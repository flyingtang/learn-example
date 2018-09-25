package main

import "fmt"

var cols = 4
var rows = 4

var arrs = [][]int{
	{1, 2, 8, 9},
	{2, 4, 9, 12},
	{4, 7, 10, 13},
	{6, 8, 11, 15}, // 这个点不能少
}

func findKey(key int) bool {

	if arrs != nil{
		var row  = 0
		var col  = cols -1

		for row < rows && col >=0 {

			if arrs[row][col] > key {
				col--
			}else if arrs[row][col] == key{
				return  true
			}else{
				row++
			}
		}
		return false
	}
}

func main() {
	key := 3
	fmt.Println(findKey(key))
}
