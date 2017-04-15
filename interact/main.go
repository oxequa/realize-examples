package main

import (
	"log"
	"os/exec"
	"bytes"
	"fmt"
	"io"
)

func main() {
	cmd := exec.Command("go", "run", "test.go")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(out.String())
	////cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "insert value")
	}()


	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	fmt.Println(out.String())
}
