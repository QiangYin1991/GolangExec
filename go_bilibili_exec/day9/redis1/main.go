package main

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	redisIp := flag.String("ip", "127.0.0.1", "redis ip")
	flag.Parse()
	ip := *redisIp + ":16379"
	c, err := redis.Dial("tcp", ip, redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("Conn redis failed. " + err.Error())
		return
	}
	defer c.Close()

	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("Get abc failed. ", err)
		return
	}
	fmt.Println("r: ", r)
}
