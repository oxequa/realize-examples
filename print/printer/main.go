package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	for i := 0; i < 10; i++ {
		fmt.Println("Tick", i)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("End")
}
