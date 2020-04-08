package model

import "github.com/qiangyin1991/tail"

type Config struct {
	LogLevel string
	LogPath  string

	ChanSize    int
	CollectConf []CollectConf
	KafkaAddr	string
}

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	Tail *tail.Tail
	Conf CollectConf
}

type TextMsg struct {
	Msg   string
	Topic string
}
