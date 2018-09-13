package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){

	path, _ := filepath.Abs(os.Args[0])
	fmt.Println("filepath ", path)
	// go run  filepath  /var/folders/sv/75hwcz6x54s4xv7fj68kw8700000gn/T/go-build942746865/b001/exe/f
	//go build filepath  /Users/xiaogang/go-learn/src/learn-example/16-filepath/app



}