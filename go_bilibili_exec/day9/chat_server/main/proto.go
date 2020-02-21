package main

import "src/go_bilibili_exec/day9/chat_server/model"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}