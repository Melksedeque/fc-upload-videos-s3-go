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

func main() {
	go contador(5)
	contador(10)
}