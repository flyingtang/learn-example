package _9_java常用算法

import (
	"fmt"
	"errors"
)

// n 个人环坐一圈，按照顺时针一次编号1 ... n, 有一个黑盒子中放置许多字条，上面有随机数字， 每人取一张，
// 游戏开始随机一个数字m，从第一个人开始，报数到m出列，并且，将手中的数字作为下一个出队数字
// 问最后剩下那个人？

// 当前的位置 随机数
// 解题思路 采用循环链表

type node struct {
	pos uint
	num uint
	next *node
}

type YosefuLink struct {
	head *node
	tail *node
}




func (ys *YosefuLink)addTailNode(num uint){

	 t := node{
	 	pos: 0,
	 	num:num,
	 }
	if ys.head == nil {
		ys.head = &t
	}
	if ys.tail == nil {
		ys.tail = &t
	}else{
		t.pos = ys.tail.pos + 1
		ys.tail.next = &t
		t.next = ys.head
		ys.tail = &t
	}
}


// popNum 第一要出队的数字
func (ys *YosefuLink)getLastOne(popNum uint) error {

	// 输入错误情况考虑
	// 0 1
	if popNum == 0 {
		return errors.New("输入出队数字要大于0")
	}

	var count = 0

	var q, p = ys.tail, ys.head
	var i uint = 0

	// 第二步 找到所有要出队的人
	for p.next != p {

		// 第一步 找到要出队的人
		i = 0
		for ; i <= popNum - 1; i++{
			q = p
			p = p.next
		}

		count++
		popNum = p.num
		fmt.Printf("找到了第 %2d 要删除的人，其位置 %2d , 手中编号 %2d \n", count, p.pos , p.num)
		// 删除之
		q.next = p.next
		p = q.next
	}


	// 第三步 输出结果
	fmt.Println("最后的家伙, ", p)
	return nil
}