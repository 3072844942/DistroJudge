package distro

import (
	"DistroJudge/api"
	"DistroJudge/compile"
	"DistroJudge/file"
	"DistroJudge/log"
	poolExecutor "DistroJudge/pool"
	snow_flake "DistroJudge/snow-flake"
	timeFormat "DistroJudge/time"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var (
	snowFlake *snow_flake.SnowFlake
	pool      *poolExecutor.Pool
	workDir   string
	cluster   *Cluster
)

type DistroConfig struct {
	WorkDir string `yaml:"work-dir"`

	Pool struct {
		MaxPoolSize uint64 `yaml:"max-pool-size"`
	} `yaml:"pool"`

	ClusterConfig ClusterConfig `yaml:"cluster"`
}

type Server struct {
	api.UnimplementedDistroServerServer
}

func NewServer(c *DistroConfig) (*Server, error) {
	var err error
	// 判题池初始化
	pool, err = poolExecutor.NewPool(c.Pool.MaxPoolSize)
	if err != nil {
		panic(fmt.Sprintf("pool executor err. err: %v", err))
	}
	// 雪花算法
	snowFlake, _ = snow_flake.GetSnowFlak(int64(rand.Intn(32)), int64(rand.Intn(32)))
	// 工作目录
	workDir = c.WorkDir
	// 集群化配置
	cluster, err = newCluster(&c.ClusterConfig, c)
	if err != nil {
		panic(fmt.Sprintf("cluster init err. err: %v", err))
	}

	return &Server{}, nil
}

// Heart /* 心跳检测, 返回当前节点的所有信息
func (d *Server) Heart(context.Context, *api.Ping) (*api.Pong, error) {
	cluster.Update()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &api.Pong{
		Cpu:                uint64(runtime.NumCPU()),
		MemoryAlloc:        m.Alloc,
		TotalAlloc:         m.TotalAlloc,
		Sys:                m.Sys,
		NumGC:              m.NumGC,
		WorkDir:            workDir,
		ActiveCount:        pool.GetRunningWorkers(),
		CompletedTaskCount: pool.GetCompileTaskCount(),
		WaitCount:          pool.GetRWaitingWorkers(),
		MaxPoolSize:        pool.GetCap(),
		Time:               time.Now().UnixMilli(),
		Status:             cluster.status,
	}, nil
}

// Join /* 加入当前节点
func (d *Server) Join(_ context.Context, node *api.Node) (*api.Cluster, error) {
	if api.Status_Leading == cluster.status {
		// 当前节点是主节点
		err := cluster.Join(node)
		if err != nil {
			return nil, err
		}
	}

	return cluster.Cluster(), nil
}

// Election /* 接受其他节点的选举策略
func (d *Server) Election(context.Context, *api.Ping) (*api.Cluster, error) {
	// 已经在选举中
	if api.Status_Looking == cluster.status {
		// 不允许被抢占, 除非从节点监控中发现了更早版本的节点
		return nil, errors.New(cluster.lastTimeStamp.Format(timeFormat.UtcTimeLayout))
	}

	// 接受其他节点的选举策略
	cluster.status = api.Status_Observing
	return cluster.Cluster(), nil
}

// Candidate /* 返回集群状态
func (d *Server) Candidate(context.Context, *api.Ping) (*api.Cluster, error) {
	return cluster.Cluster(), nil
}

// Victory /* 宣誓主权
func (d *Server) Victory(_ context.Context, node *api.Node) (*api.Node, error) {
	cluster.status = api.Status_Following
	cluster.checkMaster(node.Ip+":"+node.Port, cluster.ip, cluster.port, pool.GetCap())

	return &api.Node{
		Ip:     cluster.ip,
		Port:   cluster.port,
		Weight: pool.GetCap(),
		Status: cluster.status,
	}, nil
}

// Modify /* 更新节点信息
func (d *Server) Modify(c context.Context, distro *api.Distro) (*api.Pong, error) {
	if distro.MaxPoolSize > 0 {
		_ = pool.SetCap(distro.MaxPoolSize)
	}

	return d.Heart(c, &api.Ping{})
}

// Execute /* 执行任务
func (d *Server) Execute(c context.Context, task *api.Task) (*api.ACK, error) {
	client := cluster.Pop()
	if client != nil {
		return client.Execute(c, task)
	}
	return d.Execute(c, task)
}

// execute // 子任务执行
func (d *Server) execute(c context.Context, task *api.Task) (*api.ACK, error) {
	cluster.Update()
	comp := compile.Core{}

	dir := file.Path(workDir + "/" + task.Id)
	path, err := comp.Compile(task.Code, task.Type, dir)
	if err != nil {
		return nil, err
	}

	t := &poolExecutor.Task{
		Handler: func(v ...any) {
			run, err := comp.Run(v[0].(context.Context), v[1].(string), v[2].(api.Language), v[3].(string), v[4].(uint64), v[5].(uint64))
			if err != nil {
				log.Errorf("judge err. err: %v", err)
			}

			serverClient, err := cluster.FindClient(v[6].(string), v[7].(string))
			if err != nil {
				log.Errorf("respone %+v err. err: %v", run, err)
				return
			}

			id, _ := snowFlake.NextId()
			out, _ := file.Read(run.OutPath)
			_, _ = serverClient.Caller(c, &api.Result{
				Id:      strconv.FormatInt(id, 10),
				Out:     out,
				CpuTime: run.Time,
				Memory:  run.Memory,
			})
		},
		Params: []any{context.Background(), path, task.Type, task.In, task.CpuTime, task.Memory, task.SourceIp, task.SourcePort},
	}

	err = pool.Put(t)
	return &api.ACK{
		Id: task.Id,
	}, err
}
