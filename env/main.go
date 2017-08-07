package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("var1"))
	fmt.Println(os.Getenv("var2"))
	fmt.Println(os.Getenv("var3"))
	fmt.Println(os.Getenv("var4"))
}
