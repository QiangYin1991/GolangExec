package main

import "fmt"

func calc(intChan chan int, resultChan chan int, exitChan chan bool) {
	var val int
	for val = range intChan {
		flag := true
		for i := 2; i*i <= val; i++ {
			if val%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resultChan <- val
		}
	}

	fmt.Println(val, " exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)
	resultChan := make(chan int, 100)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 2; i < 100; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	for i := 0; i < 8; i++ {
		<-exitChan
	}

	close(resultChan)
	for v := range resultChan {
		fmt.Println("res: ", v)
	}
}
