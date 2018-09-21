package main

import (
	"errors"
	"fmt"
)

type LinkStudent struct {
	*Student
	next *LinkStudent
}

func NewLinkStudent() (*LinkStudent){
	return new(LinkStudent)
}

func (ls *LinkStudent)insert(stu *Student, keyName string) error{

	// 找位置
	if ls.IsEmpty() {
		return errors.New("链表为空，无法插入")
	}

	var pre *LinkStudent
	for th := ls; th.next != nil; th = th.next{

		if th.next.Name == keyName {
			pre = th
			break
		}

	}
	if pre == nil {
		return errors.New("未找到要插入的位置")
	}
 	// 分配内存 // 插入
	t := new(LinkStudent)
	t.Student = stu
	t.next = pre.next
	pre.next = t
	return  nil
}

func (ls *LinkStudent)appendEnd(stu *Student){

	// 找打末尾
	th := ls
	for th.next != nil {
		th = th.next
	}
	t := new(LinkStudent)
	t.Student = stu
	t.next = nil
	th.next = t
}

func (ls *LinkStudent)travel(){

	// 遍历
	for th:= ls.next;th != nil; th = th.next {
		fmt.Println(th.Student)
	}

}

func (ls *LinkStudent)IsEmpty() bool {

	if ls.next == nil {
		return  true
	}else{
		return false
	}
}