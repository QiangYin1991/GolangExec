package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"src/go_bilibili_exec/day11/logagent/tailf"
)

func main() {
	filename := "./conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n")
		panic("load conf failed")
		return
	}
	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("load conf succ, config:%v", appConfig)
	err = tailf.InitTail(appConfig.CollectConf, appConfig.ChanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("init tailf succ")

	logs.Debug("initialize success")

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("program exit")
}
