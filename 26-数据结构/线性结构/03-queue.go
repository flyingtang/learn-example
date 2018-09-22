package 线性结构

import (
	"errors"
	"fmt"
)

type QueueLine struct {

	Students []*Student
	Head, tail int
}


func NewQueueLine() *QueueLine{
	return &QueueLine{
		Students: make([]*Student, MaxLength),
	}
}


// 非环形队列
func  (ql *QueueLine)IntQueue(stu *Student) error{

	if ql.tail >= MaxLength {
		return errors.New("队列已满")
	}

	// 入队 +1
	ql.Students[ql.tail] = stu
	ql.tail++
	return  nil
}


func (ql *QueueLine)OutQueue()(*Student, error){

	if ql.Head == ql.tail {
		return nil, errors.New("队列已经空了")
	}
	stu := ql.Students[ql.Head]
	ql.Head++
	return  stu, nil
}

// 环形队列 10
//11 + 5

func (ql *QueueLine)IsFull() bool {
	if (ql.tail + 1) % MaxLength == ql.Head {
		return true
	}else {
		return false
	}
}

func (ql *QueueLine)IsEmpty() bool {

	if ql.tail == ql.Head {
		return  true
	}else{
		return false
	}

}

func  (ql *QueueLine)IntLoopQueue(stu *Student) error{

	if ql.IsFull() {
		return errors.New("队列已满")
	}

	// 入队 +1
	ql.Students[ql.tail] = stu
	ql.tail = (ql.tail + 1) % MaxLength
	//fmt.Println(ql, "===")
	return  nil
}


func (ql *QueueLine)OutLoopQueue()(*Student, error){

	if ql.IsEmpty() {
		return nil, errors.New("队列已经空了")
	}
	stu := ql.Students[ql.Head]
	ql.Head = (ql.Head +1 ) % MaxLength
	return  stu, nil
}

func (ql *QueueLine)Travel(){

	fmt.Println("====================")
	h := ql.Head
	t := ql.tail
	for t != h {
		fmt.Println("==> ",ql.Students[h])
		h = (h +1) %MaxLength
	}
	fmt.Println("====================")
}