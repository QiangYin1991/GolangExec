package main

import (
	"github.com/astaxie/beego/logs"
	"src/go_bilibili_exec/day11/logagent/kafka"
	"src/go_bilibili_exec/day11/logagent/model"
	"src/go_bilibili_exec/day11/logagent/tailf"
	"time"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = send2Kafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func send2Kafka(msg *model.TextMsg) (err error) {
	//logs.Debug("read msg:%s, topic:%s", msg.Msg, msg.Topic)
	kafka.Send2Kafka(msg.Msg, msg.Topic)
	return
}