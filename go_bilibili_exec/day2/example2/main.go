package main

import "fmt"

func main() {
	var map1 map[int]int = map[int]int{}

	map1[1] = 2
	fmt.Println(map1)

	var slice1 []string
	slice1 = append(slice1, "abc")
	fmt.Println(slice1)
}