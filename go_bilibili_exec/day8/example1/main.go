package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func read(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	intChan := make(chan int)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}
