package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
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

	logs.Debug("initialize success")
	logs.Debug("load conf succ, config:%v", appConfig)
}
