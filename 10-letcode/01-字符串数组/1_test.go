package _1_字符串数组

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	s := []int{2, 7, 11, 15}
	res := TwoSum(s, 9)
	fmt.Println("res = ", res)
	if reflect.DeepEqual([]int{0, 1}, res) {
		t.Log("Pass")
	} else {
		t.Error("error")
	}
}

func TestMyAtoi(t *testing.T) {

	MyAtoi("     -42")

}

func TestLongestCommonPrefix(t *testing.T) {

	var s = []string{"flower", "flow", "flight"}
	res := LongestCommonPrefix(s)
	fmt.Println("TestLongestCommonPrefix ", res)
}
