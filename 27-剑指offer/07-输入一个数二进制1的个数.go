package main

import "fmt"

func numbebOf1(num int) int {

	var count int
	var base int= 1
	for base > 0{

		if (num & base) > 0{
			count++
		}
		base  <<= 1
	}
	return count
}

func nunmberof1_2(n int) int{

	var count = 0

	for n > 0 {
		count++
		n = (n-1) & n
	}
	return count
}

func main()  {
	//fmt.Println(numbebOf1(0))
	//fmt.Println(numbebOf1(1))
	//fmt.Println(numbebOf1(2))
	//fmt.Println(numbebOf1(15))
	fmt.Println(nunmberof1_2(15))
}