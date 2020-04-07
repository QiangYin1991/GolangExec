package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"src/github.com/pkg/errors"
	"src/go_bilibili_exec/day11/logagent/model"
)

var (
	appConfig *model.Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var cc model.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
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

	appConfig = &model.Config{}

	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}
	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs"
	}

	appConfig.ChanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		fmt.Printf("load logs chan_size failed, err:%v\n", err)
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return
}
