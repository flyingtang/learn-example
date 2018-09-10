package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(3 * time.Second)

	go func() {
		<-t.C
		fmt.Println("t is expired")
	}()

	// t.Stop() // 这样不好
	t.Reset(0)
	fmt.Println("len ", len(t.C))
	time.Sleep(10 * time.Second)
}
