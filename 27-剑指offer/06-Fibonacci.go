package main

import "fmt"

func fabonacci_1(n int) int {

	if n <=0 {
		return  0
	}

	if n == 1 {
		return  1
	}

	return fabonacci_1(n-1) + fabonacci_1(n-2)
}

func fabonacci_2(n int) int {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var first, second, res = 0, 1, 0

	for i := 2; i <= n; i++{
		res = first + second
		first = second
		second = res
	}
	return  res
}

func main() {
	fmt.Println(fabonacci_1(4))
	fmt.Println(fabonacci_2(4))
}
