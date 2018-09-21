package main

import (
	"fmt"
)

//实现两个go轮流输出：A1B2C3.....Z26

func main() {

	// 两个数组 A -Z 1-26
	var arr1  = make([]int, 26) // 数字
	var arr2 = make([]string, 26) // 字母
	for i:= 0;i < 26; i++ {
		arr1[i] = i+1
		arr2[i] = string(uint8(65+i))
	}
	fmt.Println(arr1, arr2)

	// 循环输出 采用channel
	c1 := make(chan int)
	c2 := make(chan bool)
	go func(c  chan int){
		c2 <- true
		for{
			s := <-c
			fmt.Print(s)
			c2 <- true
		}
	}(c1)



	for i:=0;i<len(arr2);i++{
		<-c2
		fmt.Print(arr2[i])
		c1 <- arr1[i]
	}
}
