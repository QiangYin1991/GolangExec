package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"src/github.com/pkg/errors"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath	 string

	CollectConf []CollectConf
}

type CollectConf struct {
	logPath string
	topic 	string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc CollectConf
	cc.logPath = conf.String("collect::log_path")
	if len(cc.logPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}
	cc.topic = conf.String("collect::topic")
	if len(cc.topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}

	appConfig.CollectConf = append(appConfig.CollectConf, cc)
	return
}

func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	appConfig = &Config{}

	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return
}