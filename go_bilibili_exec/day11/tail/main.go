package main

import (
	"fmt"
	"github.com/qiangyin1991/tail"
	"time"
)

func main() {
	filename := "my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		Location:    nil,
		ReOpen:      false,
		MustExist:   false,
		Poll:        false,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      false,
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
			break
		}
		fmt.Println("msg: ", msg.Text)
	}
	//for line := range tails.Lines {
	//	fmt.Println(line.Text)
	//}
	//err = tails.Wait()
	//if err != nil {
	//	fmt.Println(err)
	//}
}
