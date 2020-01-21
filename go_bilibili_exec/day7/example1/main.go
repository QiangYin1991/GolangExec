package main

import (
	"fmt"
	"math/rand"
	"os"
	. "src/go_bilibili_exec/day7/example1/balance"
	"time"
)

func main() {
	var insts []*Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var balanceName string
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
