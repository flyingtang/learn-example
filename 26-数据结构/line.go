package main

import (
	"errors"
	"fmt"
)

//const MaxLength = 100


type Student struct {
	Name string
	Age uint8
}

type StudentList struct {
	Students []*Student
	Length int
}


func NewStudentList() *StudentList{

	s := new(StudentList)
	s.Students = make([]*Student, MaxLength)
	s.Length = 0 // 可以不用写
	return s
}


// n 第几个位置 1 开始
func (s *StudentList)Insert(n int, stu *Student) error{

	// 链表是否已经满了
	if s.Length >= MaxLength {
		return errors.New("线性表已满")
	}

	// n 是否越界  有数据/无数据
	if n < 1 || n > s.Length {
		return errors.New("线性表位置无效")
	}
	// 0 n s.Length
	for i := s.Length; i >= n ; i--{
		s.Students[i] = s.Students[i - 1]
	}
	s.Students[n -1] = stu
	s.Length ++
	return  nil
}

func (s *StudentList)Append(stu *Student) error{

	if s.Length >= MaxLength {
		return  errors.New("线性表已满")
	}
	s.Students[s.Length] = stu
	s.Length ++
	return nil
}

func (s *StudentList)travel(){
	for key, value :=range s.Students{
		if value == nil {
			continue
		}
		fmt.Println("key: ", key, "value: ", value)
	}
}
