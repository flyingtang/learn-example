package main

import (
	"fmt"
	"os/exec"
)

func main() {
	s := "go build ../main.go"
	cmd := exec.Command("sh", "-c", s)
	// err := cmd.Start()
	data, err := cmd.Output()

	if err != nil {
		fmt.Println("err= ", err)
	} else {
		fmt.Println("data ", string(data))
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("err= ", err)
	}
}
