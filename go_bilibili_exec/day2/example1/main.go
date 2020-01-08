package main

import (
	"fmt"
	"src/go_bilibili_exec/day2/example1/pack1"
	"src/go_bilibili_exec/day2/example1/pack2"
)

func modify(a *int) {
	fmt.Printf("modify &a: %v, a: %v, *a: %d\n", &a, a, *a)
	a = new(int)
	*a = 10
	fmt.Printf("modify &a: %v, a: %v, *a: %d\n", &a, a, *a)
}

func main() {
	fmt.Println("This is main")
	fmt.Println("pack1 Str1: ", pack1.Str1)
	fmt.Println("pack2 Str1: ", pack2.Str1)
	a := 5
	fmt.Printf("&a: %v, a: %d\n", &a, a)
	modify(&a)
	fmt.Printf("&a: %v, a: %d\n", &a, a)
	p := &a
	*p = 5
	fmt.Printf("&p: %v, p: %v, *p: %v\n", &p, p, *p)
	modify(p)
	fmt.Printf("&p: %v, p: %v, *p: %v\n", &p, p, *p)
}
