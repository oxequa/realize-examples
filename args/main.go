package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("There are no arguments")
	} else {
		fmt.Println("Args",len(os.Args))
		for _, val := range os.Args[1:] {
			fmt.Println(val)
		}
	}
}