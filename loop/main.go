package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)
	for i := 0; i < 1280; i++ {
		for i := 0; i < 768; i++ {
		}
	}
	time.Now().After(t)
	fmt.Println(time.Now())
	fmt.Println(time.Since(t))
}
