package config

import (
	"DistroJudge/distro"
	"DistroJudge/log"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DistroConfig distro.DistroConfig `yaml:"distro"`
	DbLogConfig  log.DbLogConfig     `yaml:"log"`
}

func MustLoad(c *Config, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}
}
