package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
		}
	}
	fmt.Println("You lose on")
	time.Now().After(t)

	fmt.Println(time.Now())
	fmt.Println(time.Since(t))
}
