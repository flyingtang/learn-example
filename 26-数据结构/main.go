package main

import "fmt"

const MaxLength = 3

func main()  {


 	//ss := 	NewStudentList()
	//if err := ss.Insert(1, &Student{"xiaogang", 22}); err != nil {
	//	fmt.Println(err.Error())
	//}
	//ss.Append(&Student{"xiaogang2", 11})
 	//ss.travel()
	//
	//
	//head := NewLinkStudent()
	//head.appendEnd(&Student{"xg1", 11})
	//head.appendEnd(&Student{"xg2", 11})
	//head.appendEnd(&Student{"xg3", 11})
	//
	//head.travel()
	//head.insert(&Student{"xg4", 11}, "xg2")
	//
	//fmt.Println("\n")
	//
	//head.travel()
	//fmt.Println("\n")
	//
	//nsl := NewStactLine()
	//fmt.Println(nsl.Push(&Student{"xxx1", 11}))
	//fmt.Println(nsl.Push(&Student{"xxx2", 11}))
	//
	//fmt.Println(nsl.Pop())
	//fmt.Println(nsl.Pop())
	//r,e := nsl.Pop()
	//fmt.Println(r, e, reflect.TypeOf(e))
	//fmt.Println("\n")
	//
	//nql := NewQueueLine()
	//nql.IntQueue(&Student{"ooo1", 11})
	//nql.IntQueue(&Student{"ooo1", 11})
	//
	//fmt.Println(nql.OutQueue())
	//fmt.Println(nql.OutQueue())
	//fmt.Println(nql.OutQueue())
	//fmt.Println("\n")

	nql := NewQueueLine()
	fmt.Println(nql.IntLoopQueue(&Student{"ooo1", 11}))
	fmt.Println(nql.IntLoopQueue(&Student{"ooo2", 11}))
	fmt.Println(nql.IntLoopQueue(&Student{"ooo3", 11}))

	nql.Travel()
	fmt.Println("\n")
	fmt.Println(nql.OutLoopQueue())
	nql.Travel()
	fmt.Println("\n")
	fmt.Println(nql.IntLoopQueue(&Student{"ooo3", 11}))
	nql.Travel()
	fmt.Println("\n")
	fmt.Println(nql.OutLoopQueue())
	fmt.Println(nql.OutLoopQueue())
	nql.Travel()
	fmt.Println("\n")
	//fmt.Println("\n")

}
