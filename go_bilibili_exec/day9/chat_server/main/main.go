package main

import (
	"flag"
	"time"
)

func main() {
	redisIP := flag.String("ip", "127.0.0.1", "redis ip")
	ip := *redisIP + ":16379"
	initRedis(ip, 16, 1024, time.Second * 300)
	initUserMgr()
	runServer("0.0.0.0:10000")
}
