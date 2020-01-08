package pack1

import "fmt"

var Str1 string = "pack1 string"

func init() {
	Str1 = "pack1 string init"
	fmt.Println("Str1: " + Str1)
}
