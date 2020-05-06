package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/qiangyin1991/tail"
	"src/go_bilibili_exec/day11/logagent/model"
	"time"
)

type TailObjMgr struct {
	tails   []*model.TailObj
	msgChan chan *model.TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *model.TextMsg) {
	msg = <-tailObjMgr.msgChan
	fmt.Println("msg: ", msg.Msg)
	return
}

func InitTail(cConfig []model.CollectConf, chanSize int) (err error) {
	if len(cConfig) == 0 {
		err = fmt.Errorf("invalid config for log collect, conf:%v", cConfig)
		return
	}

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *model.TextMsg, chanSize),
	}
	for _, v := range cConfig {
		obj := &model.TailObj{
			Conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			MustExist: false,
			Poll:      true,
			Follow:    true,
		})
		if errTail != nil {
			err = errTail
			return
		}
		obj.Tail = tails
		tailObjMgr.tails = append(tailObjMgr.tails, obj)

		go readFromTail(obj)
	}
	return
}

func readFromTail(tailObj *model.TailObj) {
	for true {
		msg, ok := <-tailObj.Tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.Tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		testMsg := &model.TextMsg{
			Topic: tailObj.Conf.Topic,
			Msg:   msg.Text,
		}
		tailObjMgr.msgChan <- testMsg
	}
}
