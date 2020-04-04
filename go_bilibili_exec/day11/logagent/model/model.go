package model

import "github.com/qiangyin1991/tail"

type Config struct {
	LogLevel string
	LogPath  string

	CollectConf []CollectConf
}

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	Tail *tail.Tail
	Conf CollectConf
}