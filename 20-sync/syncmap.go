package main

import (
	"sync"
	"fmt"
)

func main(){

	m := sync.Map{}

	m.Store("name", "xiaogang")

	m.Store(1, 5)

	fmt.Println( m.Load(1))
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return  true
	})
}