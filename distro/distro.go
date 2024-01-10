package distro

import (
	"DistroJudge/api"
	"DistroJudge/log"
	poolExecutor "DistroJudge/pool"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"runtime"
	"time"
)

var (
	pool *poolExecutor.Pool
)

type DistroConfig struct {
	Port int `yaml:"port"`
	Pool struct {
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

	return &Server{}, nil
}

func (d *Server) Heart(c context.Context, ping *api.Ping) (*api.Pong, error) {
	if ping.MaxPoolSize != 0 {
		logger.Infof("change pool capacity. %d -> %d", pool.GetCap(), ping.MaxPoolSize)
		pool.SetCap(ping.MaxPoolSize)
	}

	//percentages, _ := cpu.Percent(1*time.Second, false)
	//sumPercentage := float64(0)
	//for _, percentage := range percentages {
	//	sumPercentage += percentage
	//}
	//cpuUsage := sumPercentage / float64(len(percentages))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &api.Pong{
		Cpu:                0,
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

func (d *Server) Execute(context.Context, *api.Task) (*api.ACK, error) {
	return &api.ACK{}, nil
}
