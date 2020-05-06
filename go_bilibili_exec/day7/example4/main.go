package main

import "fmt"

type People interface {
	//Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

func main() {
	var stu *Student = nil
	var peo People = stu
	if peo == nil {
		fmt.Println("dddddd")
	} else {
		fmt.Println("ccccccc", peo)
	}

	if live() == nil {
		fmt.Println("a ...interface{}")
	} else {
		fmt.Printf("BBBBBBBB %v\n", live())
	}

	ret := testDefer()
	fmt.Println(ret)
}

func testDefer() (ret int) {
	defer func() {
		ret = 2
	}()
	fmt.Println(ret)
	return 1
}
