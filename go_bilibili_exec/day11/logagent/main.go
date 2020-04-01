package main

import "fmt"

func main() {
	err := loadConf()
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n")
	}
}
