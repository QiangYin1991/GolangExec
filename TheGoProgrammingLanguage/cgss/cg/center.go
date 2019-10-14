/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-14 17:37:00
 * @LastEditTime: 2019-10-14 18:42:57
 * @LastEditors: Please set LastEditors
 */
package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"TheGoProgrammingLanguage/cgss/ipc"
)

var ipc.Server = &CenterServer{}	//确认实现了Server接口

type Message struct {
	From 	string "from"
	To	 	string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string] ipc.Server
	players []*Player
	rooms []*Room
	mutex sync.RWMutex
}

func NewCenterServer struct {
	servers := make(map[string] ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers:servers, players:players}
}