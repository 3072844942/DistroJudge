package distro

import (
	"DistroJudge/api"
	"DistroJudge/compile"
	"DistroJudge/file"
	"DistroJudge/log"
	poolExecutor "DistroJudge/pool"
	snow_flake "DistroJudge/snow-flake"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var (
	snowFlake *snow_flake.SnowFlake
	pool      *poolExecutor.Pool
	client    map[string]api.DistroServerClient
	workDir   string
)

type DistroConfig struct {
	Port    int    `yaml:"port"`
	WorkDir string `yaml:"work-dir"`
	Pool    struct {
		MaxPoolSize uint64 `yaml:"max-pool-size"`
	}
}

type Server struct {
	api.UnimplementedDistroServerServer
}

func NewServer(c *DistroConfig) (*Server, error) {
	var err error

	pool, err = poolExecutor.NewPool(c.Pool.MaxPoolSize)
	if err != nil {
		log.Errorf("pool executor err. err: %v", err)
	}

	snowFlake, _ = snow_flake.GetSnowFlak(int64(rand.Intn(32)), int64(rand.Intn(32)))

	workDir = c.WorkDir

	return &Server{}, nil
}

func (d *Server) Heart(c context.Context, ping *api.Ping) (*api.Pong, error) {
	if ping.MaxPoolSize != 0 {
		log.Infof("change pool capacity. %d -> %d", pool.GetCap(), ping.MaxPoolSize)
		pool.SetCap(ping.MaxPoolSize)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &api.Pong{
		Cpu:                uint64(runtime.NumCPU()),
		MemoryAlloc:        m.Alloc,
		TotalAlloc:         m.TotalAlloc,
		Sys:                m.Sys,
		NumGC:              m.NumGC,
		ActiveCount:        pool.GetRunningWorkers(),
		CompletedTaskCount: pool.GetCompileTaskCount(),
		WaitCount:          pool.GetRWaitingWorkers(),
		MaxPoolSize:        pool.GetCap(),
		Time:               time.Now().UnixMilli(),
	}, nil
}

func (d *Server) Execute(c context.Context, task *api.Task) (*api.ACK, error) {
	comp := compile.Core{}

	dir := file.Path(workDir + "/" + task.Id)
	path, err := comp.Compile(task.Code, task.Type, dir)
	if err != nil {
		return nil, err
	}

	t := &poolExecutor.Task{
		Handler: func(v ...any) {
			run, err := comp.Run(v[0].(context.Context), v[1].(string), v[2].(api.Task_Language), v[3].(string), v[4].(uint64), v[5].(uint64))
			if err != nil {
				log.Errorf("judge err. err: %v", err)
			}

			ip, port := v[6].(string), v[7].(string)
			addr := ip + ":" + port

			if _, st := client[addr]; !st {
				conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					log.Errorf("connect %s:%s err. err: %v", ip, port, err)
				}

				client[addr] = api.NewDistroServerClient(conn)
			}
			serverClient := client[addr]

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
