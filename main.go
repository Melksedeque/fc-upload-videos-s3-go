package main

import (
	"fmt"
	"time"
)

func contador(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	data := make(chan int)
	go worker(1, data)
	go worker(2, data)

	for i := 0; i < 10; i++ {
		go worker(i, data)
	}
	
	for i := 0; i < 100; i++ {
		data <- i
	}

	time.Sleep(time.Second * 10)
}