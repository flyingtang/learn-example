package interview

import (
	"container/list"
	"fmt"
)

// 构成顺序数组
func make_array(n int) *list.List {
	array := list.New()
	for i := 1; i <= n; i++ {
		array.PushBack(i)
	}
	return array
}

// 生成中间数组
func count_mid_array(array *list.List) *list.List {
	mid_array := list.New()
	for array.Len() > 1 {
		t1 := array.Front()
		mid_array.PushBack(t1.Value)
		array.Remove(t1)
		t2 := array.Front()
		array.Remove(t2)
		array.PushBack(t2.Value)
	}
	mid_array.PushBack(array.Front().Value)
	return mid_array
}

// 转换中间组到最终结果
func trans_array(n int, new_array *list.List) *list.List {
	fin_array := list.New()
	for i := 1; i <= n; i++ {
		idx := 0
		for e := new_array.Front(); e != nil; e = e.Next() {
			idx++
			if e.Value == i {
				fin_array.PushBack(idx)
			}
		}
	}
	return fin_array
}

// 打印list函数
func PrintList(l *list.List) {
	if l != nil {
		fmt.Printf("[")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d", e.Value)
			if e.Next() != nil {
				fmt.Print(", ")
			}
		}
		fmt.Printf("]\n")
	}
}
// 约瑟夫环入口
//func main() {
//	fmt.Println("----====joseph run====----")
//	var n int = 12
//	array := make_array(n)
//	PrintList(array)
//	mid_array := count_mid_array(array)
//	PrintList(mid_array)
//	fin_array := trans_array(n, mid_array)
//	PrintList(fin_array)
//}