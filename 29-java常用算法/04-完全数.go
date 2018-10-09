package _9_java常用算法

import (
	"fmt"
)

// 完全数 就是所有真因子之和等于这个数
// 真因子 就是除本身之外的所有因子
// 因子 就是所有能被整除的数 例如6 => 6 3 2 1

// 找出这个num 之内的所有数的完全数
func FindAllNumber(num int) {

	if num < 1 {
		return
	}
	var temp = make([]int, num)
	var sum, count = 0, 0
	for i := 1; i <= num; i++ {
		sum = 0
		count = 0

		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				sum += j
				temp[count] = j
				count++
			}
		}

		if sum == i {
			fmt.Printf("current num  %d is complete number: \n", i)
			for t := 0; t < count; t++ {
				fmt.Printf("%d \t", temp[t])
			}
			fmt.Println()
		}
	}
}