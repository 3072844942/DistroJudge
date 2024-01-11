package main

import (
	"DistroJudge/api"
	"DistroJudge/config"
	"DistroJudge/distro"
	"DistroJudge/log"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	configFile = flag.String("f", "etc/dev.yml", "the config file")
)

func main() {
	flag.Parse()

	var c config.Config
	config.MustLoad(&c, *configFile)

	// 日志选项
	client := log.NewClient(&c.DbLogConfig)
	defer client.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.DistroConfig.Port))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	d, err := distro.NewServer(&c.DistroConfig)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	api.RegisterDistroServerServer(s, d)
	log.Infof("server listening at %v", lis.Addr())

	//prometheus监控
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()

	if err = s.Serve(lis); err != nil {
		log.Errorf("failed to server: %v", err)
	}
}
