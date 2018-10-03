package main

import (
	"fmt"
)

var graph map[string][]string


// 不考虑线程安全问题
func NewGraph(){

	graph = make(map[string][]string)
	graph["A"] = []string{
		"B","C",
	}
	graph["B"] = []string{
		"A","C","D",
	}
	graph["C"] = []string{
		"A","B","D","E",
	}
	graph["D"] = []string{
		"B","C","E","F",
	}
	graph["E"] = []string{
		"C","D",
	}
	graph["F"] = []string{
		"D",
	}
}


func IsExists(source []string, key string) bool{
	length := len(source)
	for i:=0;i< length ;i++  {
		if source[i] == key {
			return true
		}
	}
	return  false
}

// 简单的广度优先遍历
func BFS_v1(s string){
	queue := NewQueue()
	queue.Push(s)

	var seen = []string{s,}

	for  length := len(queue.data);length > 0; length = len(queue.data) {
		if vet, err := queue.Pop(); err != nil {
			panic("程序数据出现异常")
		}else{
			vetstr := vet.(string)
			nodes := graph[vetstr]
			nLength := len(nodes)
			for i:= 0 ; i < nLength; i++{
				if !IsExists(seen, nodes[i]){
					seen = append(seen, nodes[i])
					queue.data = append(queue.data, nodes[i])
				}
			}
			fmt.Println(" ===> ",vet)
		}

	}
}



// 简单的深度优先遍历
func DFS_v1(s string) {
	queue := NewStack()
	queue.Push(s)

	var seen = []string{s,}
	for  length := len(queue.data);length > 0; length = len(queue.data) {
		if vet, err := queue.Pop(); err != nil {
			panic("程序数据出现异常")
		}else{
			vetstr := vet.(string)
			nodes := graph[vetstr]
			nLength := len(nodes)
			for i:= 0 ; i < nLength; i++{
				if !IsExists(seen, nodes[i]){
					seen = append(seen, nodes[i])
					queue.data = append(queue.data, nodes[i])

				}
			}
			fmt.Println(" ===> ",vet)
		}

	}

}
//
func DFS_v2(s string)map[string]string {
	queue := NewStack()
	queue.Push(s)

	var seen = []string{s,}
	var parent = map[string]string{
		s: "NONE",
	}
	for  length := len(queue.data);length > 0; length = len(queue.data) {
		if vet, err := queue.Pop(); err != nil {
			panic("程序数据出现异常")
		}else{
			vetstr := vet.(string)
			nodes := graph[vetstr]
			nLength := len(nodes)
			for i:= 0 ; i < nLength; i++{
				if !IsExists(seen, nodes[i]){
					seen = append(seen, nodes[i])
					queue.data = append(queue.data, nodes[i])
					parent[nodes[i]] = vetstr
				}
			}
			fmt.Println(" ===> ",vet)
		}

	}
	return parent
}




//func main(){
//	NewGraph()
//	//fmt.Println("BFS")
//	//BFS("A")
//	//fmt.Println("DFS")
//	parent := DFS_v2("E")
//	//for key, value := range parent {
//	//	fmt.Println(key, "===> ", value)
//	//}
//
//	var v = "B"
//	for v != "NONE" {
//		fmt.Println(v)
//		v = parent[v]
//	}
//
//}