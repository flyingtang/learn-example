package main

import "errors"

type Stack struct {
	data []interface{}
}

func NewStack() *Stack{
	return  &Stack{
		data:make([]interface{}, 0),
	}
}

func (q *Stack)IsEmpty()  bool {
	if len(q.data) == 0 {
		return  true
	}else {
		return false
	}
}

func (q *Stack)Push(s string){

	q.data = append(q.data, s)
}
func (q *Stack)Pop() (s interface{}, err error){
	if q.IsEmpty() {
		err = errors.New("Empty Stack")
		return
	}
	length := len(q.data)
	s = q.data[length - 1]
	q.data = q.data[:length - 1]
	return
}