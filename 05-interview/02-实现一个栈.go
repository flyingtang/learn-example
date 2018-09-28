package interview

import (
	"errors"
	"fmt"
	"sync"
)

// 实现一个栈， 然后大量的goroutine 往里面写数据，还有出栈

type Stack struct {
	data []int
	top  int
}

var l = sync.Mutex{}

const MaxLength = 100

// 不考虑边界输入边界什么的特殊情况
//版本1 用锁来实现 的数组来实现
// 缺点，栈是有限的 100

func NewStack() *Stack {
	return &Stack{
		data: make([]int, MaxLength),
		top:  0,
	}
}

func (this *Stack) pop() (int, error) {
	l.Lock()
	defer l.Unlock()

	t := this.top
	if t == 0 {
		return 0, errors.New("栈已经空了")
	}

	value := this.data[t-1]
	this.data[t-1] = 0
	this.top--
	return value, nil
}

func (this *Stack) push(value int) error {
	l.Lock()
	defer l.Unlock()

	if this.top == MaxLength {
		return errors.New("栈已经满了")
	}
	this.data[this.top] = value
	this.top++
	return nil
}

// 版本2
type chData struct {
	dir   string
	value int
}

type Stack_2 struct {
	data []int
	in   chan chData
	out  chan int
}

const chanSize = 10

// 不考虑边界输入边界什么的特殊情况
func NewStack_2() (s2 *Stack_2) {

	s2 = &Stack_2{
		data: make([]int, 0),
		in:   make(chan chData, chanSize),
		out:  make(chan int, chanSize),
	}
	go func() {
		for {
			select {
			case cd := <-s2.in:
				{
					switch cd.dir {
					case "push":
						s2.data = append(s2.data, cd.value)
					case "pop":
						s2length := len(s2.data)
						if s2length > 0{
							t := s2.data[s2length-1]
							s2.data = s2.data[:s2length-1]
							s2.out <- t
						}else{
							fmt.Println("栈已经空啦")
						}
					}
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case value := <-s2.out:
				fmt.Println("输出 ==> ", value)
			}
		}
	}()
	return s2
}

func (this *Stack_2) pop_2() {
	var cd = chData{
		dir: "pop",
	}

	this.in <- cd
	return
}

func (this *Stack_2) push_2(value int) {

	var cd = chData{
		dir:   "push",
		value: value,
	}

	this.in <- cd
	return
}
