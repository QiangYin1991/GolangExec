package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		defer fmt.Println("Close goroutine a")

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer fmt.Println("Close goroutine A")

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer fmt.Println("Close goroutine C")

		for count := 0; count < 3; count++ {
			for char := 0; char < 26; char++ {
				fmt.Printf("%d ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
