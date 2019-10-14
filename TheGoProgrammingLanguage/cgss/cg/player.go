/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-14 17:36:51
 * @LastEditTime: 2019-10-14 17:59:46
 * @LastEditors: Please set LastEditors
 */
package cg

import (
	"fmt"
)

type Player struct {
	Name string "name"
	Level int "level"
	Exp int "exp"
	Room int "room"

	mq chan *Message // 等待收取的消息
}

func NewPlayer() *Player  {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)

	return player;
}