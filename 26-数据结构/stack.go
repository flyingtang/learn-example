package main

import "errors"

//const MaxLength = 100

type StackLine struct {
	Students  []*Student

	Length int
}


func NewStactLine() *StackLine{

	return &StackLine{
		Students: make([]*Student, MaxLength),
		Length: 0,
	}
}

func (sl *StackLine)Push(stu *Student) error{

	// 栈满了
	if sl.Length >= MaxLength {
		return errors.New("栈已经满了")
	}
	// stu 为空
	if stu == nil {
		return errors.New("插入的数据无效")
	}
	//插入
	sl.Students[sl.Length] = stu
	// 长度加一
	sl.Length ++

	return  nil
}

func (sl *StackLine)Pop()( *Student,  error){

	// 栈空
	if sl.Length == 0 {
		return nil, errors.New("栈空")
	}
	// 出栈

	stu := sl.Students[sl.Length - 1]
	// 长度减一
	sl.Length --
	return stu, nil
}




