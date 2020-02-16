package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	redisIP := flag.String("ip", "127.0.0.1", "redis ip")
	flag.Parse()
	ip := *redisIP + ":16379"
	pool := &redis.Pool{
		TestOnBorrow: nil,
		MaxIdle:      16,
		MaxActive:    1024,
		IdleTimeout:  300,
		Wait:         false,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", ip, redis.DialPassword("123456"))
		},
	}

	c := pool.Get()
	defer c.Close()

	_, err := c.Do("HSet", "books", "abc", 120)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("Get abc failed. ", err)
		return
	}
	fmt.Println("r:", r)

	pool.Close()
}
