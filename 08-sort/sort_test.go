package sort

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	s := []int{
		15,10,2,4,5,9,
	}
	fmt.Println("排序前： ", s)
	BubbleSort(s)
	fmt.Println("排序后： ", s)
}
