package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	index    int
	value    string
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {

	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

var pqGraph map[string]map[string]int

// 不考虑线程安全问题
func NewGraph_v2() {

	pqGraph = make(map[string]map[string]int)
	pqGraph["A"] = map[string]int{
		"B": 5, "C": 1,
	}
	pqGraph["B"] = map[string]int{
		"A": 5, "C": 2, "D": 1,
	}
	pqGraph["C"] = map[string]int{
		"A": 1, "B": 2, "D": 4, "E": 8,
	}
	pqGraph["D"] = map[string]int{
		"B": 1, "C": 4, "E": 3, "F": 6,
	}
	pqGraph["E"] = map[string]int{
		"C": 8, "D": 3,
	}
	pqGraph["F"] = map[string]int{
		"D": 6,
	}
}

func InitDistance(s string) (m map[string]int) {

	m = make(map[string]int)
	m[s] = 0

	for key := range pqGraph {
		if key != s {
			m[key] = math.MaxInt32
		}
	}

	return m
}

// 根据权重找到最短路径
func Dijkstra(s string) (map[string]string, map[string]int) {

	var queue = make(PriorityQueue, 0)
	heap.Push(&queue, &Item{value: s, priority: 0})
	heap.Init(&queue)

	var seen = make([]string, 0)
	var parent = map[string]string{
		s: "NONE",
	}
	distance := InitDistance(s)

	for length := len(queue); length > 0; length = len(queue) {

		vet := heap.Pop(&queue).(*Item)
		vetex := vet.value
		dis := vet.priority
		nodes := pqGraph[vetex]
		seen = append(seen, vetex)

		for key := range nodes {
			if !IsExists(seen, key) {
				if dis+pqGraph[vetex][key] < distance[key] {
					parent[key] = vetex
					distance[key] = dis + pqGraph[vetex][key]
					heap.Push(&queue, &Item{value: key, priority: dis + pqGraph[vetex][key]})

				}

			}
		}
	}
	return parent, distance
}

func main() {

	NewGraph_v2()
	p, d := Dijkstra("A")
	fmt.Println(p)
	fmt.Println(d)
}

func printQueue(queue PriorityQueue){
	for key, value := range queue{
		fmt.Println("key ", key, " value ",  value)
	}
	fmt.Println("\n")
}