package main

import "fmt"
import _ "learn-example/17-init/testpack"

//func init(){
//	fmt.Println("main init 1")
//}
//func init(){
//	fmt.Println("main init 2")
//}

func main(){
	fmt.Println("1111")
	fmt.Println("2222")
}

//
//testpack init 1
//testpack init 2
//main init 1
//main init 2
//1111
//2222

func init(){
	fmt.Println("main init 1")
}
func init(){
	fmt.Println("main init 2")
}