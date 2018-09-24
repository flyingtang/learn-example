package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name string
}

type ChainTree struct {

	Student Student
	left *ChainTree
	right *ChainTree
}

func NewNode(name string)(c *ChainTree){

	c = new(ChainTree)
	c.Student = Student{Name:name}
	c.left = nil
	c.right = nil
	return c
}

//const MaxLength = 100 // 设置一个循环队列

// 层级遍历
func (ct *ChainTree)LevelTravel(){

	//  空树直接返回
	if ct == nil {
		return
	}

	ctTemp := ct
	var head,tail int
	queue := make([]*ChainTree, MaxLength)


	if ctTemp != nil {

		tail = (tail+1) % MaxLength
		queue[tail] = ctTemp
	}
	for head != tail {

		head = (head + 1) % MaxLength
		t := queue[head]
		fmt.Println("==》 ", t.Student.Name)

		if t.left != nil {
			tail = (tail+1) % MaxLength
			queue[tail] = t.left
		}

		if t.right != nil {
			tail = (tail+1) % MaxLength
			queue[tail] = t.right
		}
	}

}

// 前序遍历
func DLRTree(c *ChainTree){

	if c != nil{

		fmt.Println("==> ", c.Student.Name)
		DLRTree(c.left)
		DLRTree(c.right)
	}
}

// 中序遍历
func LDRTree(c *ChainTree){

	if c != nil{

		LDRTree(c.left)
		fmt.Println("==> ", c.Student.Name)
		LDRTree(c.right)
	}
}

// 后序遍历
func LRDTree(c *ChainTree) {
	if c != nil {

		LRDTree(c.left)
		LRDTree(c.right)
		fmt.Println("==> ", c.Student.Name)
	}
}

// 获取树的深度
func TreeDepth(c *ChainTree) int {

	if c == nil {
		return  0
	}

	left := TreeDepth(c.left)
	right := TreeDepth(c.right)

	res := int(math.Max(float64(left), float64(right)))

	return res + 1
}

// 清空链表 会不会出错
func ClearTree(c *ChainTree)  {

	// 刚开始这里写错了，导致清不掉，如果不理解go的函数值传类型，会想不明白
	if c != nil{
		ClearTree(c.left)
		ClearTree(c.right)
		c.left = nil  // 特别这里
		c.right = nil // 特别这里
		c.Student = Student{}
	}

}
//func main(){
//
//	root := NewNode("1")
//
//	rootLeft := NewNode("2")
//	rootRight := NewNode("3")
//
//	root.left = rootLeft
//	root.right = rootRight
//
//	rootLeftL := NewNode("4")
//	rootLeftR := NewNode("5")
//
//	rootLeft.left = rootLeftL
//	rootLeft.right = rootLeftR
//
//	// 树的深度
//	h := TreeDepth(root)
//	fmt.Println("树的深度： ", h, "\n")
//
//	// 层级遍历
//	fmt.Println("层级遍历: ")
//	root.LevelTravel()
//	fmt.Println("\n")
//
//	fmt.Println("前序遍历: ")
//	DLRTree(root)
//	fmt.Println("\n")
//
//	fmt.Println("中序遍历: ")
//	LDRTree(root)
//	fmt.Println("\n")
//
//	fmt.Println("后序遍历: ")
//	LRDTree(root)
//	fmt.Println("\n")
//
//
//	// 有问题
//	fmt.Println(root)
//	fmt.Println("====清理树====")
//	ClearTree(root)
//	fmt.Println(root)
//}
//
//树的深度：  3
//
//层级遍历:
//==》  1
//==》  2
//==》  3
//==》  4
//==》  5
//
//
//前序遍历:
//==>  1
//==>  2
//==>  4
//==>  5
//==>  3
//
//
//中序遍历:
//==>  4
//==>  2
//==>  5
//==>  1
//==>  3
//
//
//后序遍历:
//==>  4
//==>  5
//==>  2
//==>  3
//==>  1
//
//
//&{{1} 0xc42000a0a0 0xc42000a0c0}
//====清理树====
//&{{} <nil> <nil>}

