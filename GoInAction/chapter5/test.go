package main

import "fmt"

func main() {
	var arr = [3]int{10, 20, 30}

	for i, v := range arr {
		if i == 0 {
			arr[0] += 100
			arr[1] += 200
			arr[2] += 300
		}
		fmt.Println("i:", i, " v:", v)
	}

	for i, v := range arr[:] {
		if i == 0 {
			fmt.Println("arr:", arr[0], arr[1], arr[2])
			arr[0] += 100
			arr[1] += 200
			arr[2] += 300
			fmt.Println("arr:", arr[0], arr[1], arr[2])
		}
		fmt.Println("i:", i, " v:", v)
	}
}
