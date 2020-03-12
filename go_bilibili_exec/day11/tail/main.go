package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main()  {
	filename := "my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		Location:    nil,
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg: ", msg)
	}
}
