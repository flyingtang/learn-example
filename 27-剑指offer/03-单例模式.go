package main

import "sync"

type Student struct {
	Name string
}


var student *Student
var m = sync.Mutex{}



func NewSingleStudent() *Student{

	if student == nil{
		m.Lock()
		if student == nil {
			student = new(Student)
		}
		m.Unlock()
	}
	return student
}


