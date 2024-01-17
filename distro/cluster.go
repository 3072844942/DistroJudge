package distro

import (
	"DistroJudge/api"
	timeFormat "DistroJudge/time"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"sync"
	"time"
)

var (
	ce *Cluster
)

type ClusterConfig struct {
	Master     bool   `yaml:"master"`
	Ip         string `yaml:"ip"`
	Port       string `yaml:"port"`
	MasterAddr string `yaml:"master-addr"`
}

type Cluster struct {
	sync.Mutex
	ip   string
	port string
	// 响应端
	client map[string]api.DistroServerClient
	// 主节点配置
	status       api.Status
	masterAddr   string
	masterServer api.DistroServerClient
	// 集群, 使用加权随机
	cluster    map[string]api.DistroServerClient
	weightTree map[string]uint64
	// 从节点配置
	lastTimeStamp time.Time
}

// newCluster 创建一个集群
func newCluster(c *ClusterConfig, distro *DistroConfig) (*Cluster, error) {
	ce = &Cluster{
		ip:         c.Ip,
		port:       c.Port,
		client:     make(map[string]api.DistroServerClient),
		status:     api.Status_Looking,
		cluster:    make(map[string]api.DistroServerClient),
		weightTree: make(map[string]uint64),
	}

	if c.Master {
		ce.status = api.Status_Leading
		ce.masterAddr = c.Ip + ":" + c.Port
	}

	if !c.Master {
		ce.status = api.Status_Following
		ce.masterAddr = c.MasterAddr

		// 从节点连接主节点
		ce.checkMaster(c.MasterAddr, c.Ip, c.Port, distro.Pool.MaxPoolSize)

		// 从节点监控主节点
		go ce.monitor()
	}

	return ce, nil
}

// Pop 找到最近空闲的节点
func (c *Cluster) Pop() api.DistroServerClient {
	c.Lock()
	defer c.Unlock()

	totalWeight := int64(0)
	for _, w := range c.weightTree {
		totalWeight += int64(w)
	}

	pos := uint64(rand.Int63n(totalWeight))
	for addr, w := range c.weightTree {
		if pos < w {
			return c.cluster[addr]
		}

		pos -= w
	}
	return nil
}

// Update /* 更新版本时间戳
func (c *Cluster) Update() {
	c.lastTimeStamp = time.Now()
}

// Join /* 加入新的节点到集群
func (c *Cluster) Join(nodes ...*api.Node) error {
	for _, node := range nodes {
		client, err := c.FindClient(node.Id, node.Port)
		if err != nil {
			return err
		}
		addr := node.Ip + ":" + node.Port
		c.weightTree[addr] = node.Weight
		c.cluster[addr] = client
	}
	return nil
}

// FindClient /* 查找响应方
func (c *Cluster) FindClient(ip, port string) (api.DistroServerClient, error) {
	var err error
	addr := ip + ":" + port

	if _, st := c.client[addr]; !st {
		c.client[addr], err = c.connect(addr)
		if err != nil {
			return nil, err
		}
	}
	return c.client[addr], nil
}

// checkMaster /* 切换主节点
func (c *Cluster) checkMaster(addr, localIp, localPort string, weight uint64) {
	var err error

	c.masterAddr = addr
	c.masterServer, err = c.connect(addr)
	if err != nil {
		panic(fmt.Sprintf("connect master %v err. err: %v", ce.masterAddr, err))
	}

	joined, err := c.masterServer.Join(context.Background(), &api.Node{
		Ip:     localIp,
		Port:   localPort,
		Weight: weight,
	})
	if err != nil {
		panic(fmt.Sprintf("join cluster err. err: %v", err))
	}

	// 如果所连节点不是主节点
	if joined.MasterAddr != c.masterAddr {
		c.checkMaster(joined.MasterAddr, localIp, localPort, weight)
	}
	// 保存集群信息备用
	for _, i := range joined.Addr {
		c.weightTree[i] = 0
	}
}

// connect /* 连接distro节点
func (c *Cluster) connect(addr string) (api.DistroServerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return api.NewDistroServerClient(conn), nil
}

// clusterMonitor /* 主节点集群监控
func (c *Cluster) clusterMonitor() {
	for {
		// 主节点监控所有从节点是否活跃
		for addr, cl := range c.cluster {
			ctx, canalFunc := context.WithTimeout(context.Background(), 2*time.Second)
			_, err := cl.Heart(ctx, &api.Ping{})
			canalFunc()

			if err != nil {
				c.Lock()
				delete(c.cluster, addr)
				delete(c.weightTree, addr)
				c.Unlock()
			}
		}

		time.Sleep(3 * time.Second)
	}
}

// monitor /* 从节点监控
func (c *Cluster) monitor() {
	for {
		time.Sleep(3 * time.Second)

		now := time.Now()
		if now.Sub(c.lastTimeStamp).Seconds() <= 3 {
			continue
		}

		// 超过3s没有响应
		_, err := c.masterServer.Heart(context.Background(), &api.Ping{})
		if err == nil {
			// 主节点正常
			continue
		}

		// (如果当前节点正在选举流程中 && 10s仍没有选举完成) || 当前节点为从节点,开启选举过程
		if (c.status == api.Status_Observing && now.Sub(c.lastTimeStamp).Seconds() > 10) || c.status == api.Status_Following {
			// 开启选举过程
			c.status = api.Status_Looking
			c.election()
		}
	}
}

// Cluster /* 查询集群信息
func (c *Cluster) Cluster() *api.Cluster {
	res := &api.Cluster{
		MasterAddr: cluster.masterAddr,
		Addr:       make([]string, len(cluster.weightTree)),
		ClientAddr: make([]string, len(cluster.weightTree)),
	}
	for addr := range cluster.weightTree {
		res.Addr = append(res.Addr, addr)
	}
	for addr := range cluster.client {
		res.ClientAddr = append(res.ClientAddr, addr)
	}
	return res
}

// election /* 选举
func (c *Cluster) election() {
	var err error
	ctx := context.Background()

	// 询问所有节点, 保存集群信息
	for key := range c.weightTree {
		if _, st := c.cluster[key]; !st {
			c.cluster[key], err = c.connect(key)
			if err != nil {
				delete(c.cluster, key)
			}
		}
	}

	// 去除所有
	for addr, se := range c.cluster {
		_, err := se.Heart(ctx, &api.Ping{})
		if err != nil {
			delete(c.cluster, addr)
		}
	}

	for _, se := range c.cluster {
		cl, err := se.Election(ctx, &api.Ping{})
		if err != nil {
			// 如果该节点也在选举流程
			timeStamp, _ := time.Parse(timeFormat.UtcTimeLayout, err.Error())
			if c.lastTimeStamp.After(timeStamp) {
				// 该节点优选度高于当前节点
				// 放弃选举权, 接受安排
				c.status = api.Status_Observing
				return
			}

			// 休眠, 等待节点释放选举权
			time.Sleep(time.Second)

			// 再次抢占选举权
			cl, _ = se.Election(ctx, &api.Ping{})
		}

		// 遍历该节点保存的所有集群
		for _, key := range cl.Addr {
			if _, st := c.cluster[key]; !st {
				c.cluster[key], err = c.connect(key)
				if err != nil {
					delete(c.cluster, key)
				}
			}
		}

		// 遍历该节点保存的所有客户端集群
		for _, key := range cl.ClientAddr {
			if _, st := c.client[key]; !st {
				c.client[key], err = c.connect(key)
				if err != nil {
					delete(c.client, key)
				}
			}
		}
	}

	// 当前节点成为主节点
	c.status = api.Status_Leading
	// 宣誓主权
	for addr, se := range c.cluster {
		node, _ := se.Victory(ctx, &api.Node{
			Ip:   c.ip,
			Port: c.port,
		})

		c.weightTree[addr] = node.Weight
	}
	// 通知客户端
	info := c.Cluster()
	for _, se := range c.client {
		_, _ = se.Checkout(ctx, info)
	}
}
