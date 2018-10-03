package main

import "errors"

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue{
	return  &Queue{
		data:make([]interface{}, 0),
	}
}

func (q *Queue)IsEmpty()  bool {
	if len(q.data) == 0 {
		return  true
	}else {
		return false
	}
}

func (q *Queue)Push(s string){

	q.data = append(q.data, s)
}
func (q *Queue)Pop() (s interface{}, err error){
	if q.IsEmpty() {
		err = errors.New("Empty Queue")
		return
	}

	s = q.data[0]
	q.data = q.data[1:]
	return
}