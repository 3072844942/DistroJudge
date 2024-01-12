package distro

import (
	"DistroJudge/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

type ClusterConfig struct {
	Master bool   `yaml:"master"`
	Ip     string `yaml:"ip"`
	Port   string `yaml:"port"`
}

type Cluster struct {
	client       map[string]api.DistroServerClient
	master       bool
	masterAddr   string
	masterServer api.DistroServerClient
	cluster      map[string]api.DistroServerClient
	weightTree   map[string]int64
}

func newCluster(c *ClusterConfig) (ce *Cluster, err error) {
	ce = &Cluster{
		client:     make(map[string]api.DistroServerClient),
		master:     c.Master,
		masterAddr: c.Ip,
		cluster:    make(map[string]api.DistroServerClient),
		weightTree: make(map[string]int64),
	}
	if c.Master {
		ce.masterAddr = c.Ip + ":" + c.Port
		ce.masterServer, err = connect(ce.masterAddr)
	}
	return ce, nil
}

func (c *Cluster) Pop() api.DistroServerClient {
	totalWeight := int64(0)
	for _, w := range c.weightTree {
		totalWeight += w
	}

	pos := rand.Int63n(totalWeight)
	for addr, w := range c.weightTree {
		if pos < w {
			return c.cluster[addr]
		}

		pos -= w
	}
	return nil
}

func (c *Cluster) Join(node *api.Node) error {
	client, err := c.FindClient(node.Id, node.Port)
	if err != nil {
		return err
	}
	addr := node.Ip + ":" + node.Port
	c.weightTree[addr] = node.Weight
	c.cluster[addr] = client
	return nil
}

func (c *Cluster) FindClient(ip, port string) (api.DistroServerClient, error) {
	var err error
	addr := ip + ":" + port

	if _, st := c.client[addr]; !st {
		c.client[addr], err = connect(addr)
		if err != nil {
			return nil, err
		}
	}
	return c.client[addr], nil
}

func connect(addr string) (api.DistroServerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return api.NewDistroServerClient(conn), nil
}
