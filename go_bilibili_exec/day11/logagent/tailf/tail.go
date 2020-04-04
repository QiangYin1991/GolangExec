package tailf

import (
	"fmt"
	"github.com/qiangyin1991/tail"
	"src/go_bilibili_exec/day11/logagent/model"
)

type TailObjMgr struct {
	tails []*model.TailObj
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(cConfig []model.CollectConf) (err error) {
	if len(cConfig) == 0 {
		err = fmt.Errorf("invalid config for log collect, conf:%v", cConfig)
		return
	}

	tailObjMgr = &TailObjMgr{}
	for _, v := range cConfig {
		obj := &model.TailObj{
			Conf: model.CollectConf{},
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:      true,
			MustExist:   false,
			Poll:        true,
			Follow:      true,
		})
		if errTail != nil {
			err = errTail
			return
		}
		obj.Tail = tails
		tailObjMgr.tails = append(tailObjMgr.tails, obj)
	}
	return
}
