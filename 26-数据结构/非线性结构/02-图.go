package main


// 图 邻接矩阵的实现方式


import (
	"errors"
	"fmt"
)

const MaxLength = 100

type GraphMatrix struct {
	Data     []string // 保存数据
	gType    int      // 0 无效图 1 有向 一旦初始化，就不改变
	VNum     int      // 定点的数量
	ENum     int      // 边的数量
	EWeight  [][]int  // 边的权值
	IsTravel []int    // 遍历标志
}

// 初始化一个图
func NewGraphMatrix(GType int) (*GraphMatrix, error) {

	if GType != 0 && GType != 1 {
		return nil, errors.New("GType must 0 or 1")
	}

	// 二维数组初始化，讲究了哈
	t := make([][]int, MaxLength)
	for i := 0; i < MaxLength; i++{
		t[i] = make([]int, MaxLength)
	}
	return &GraphMatrix{
		Data:     make([]string, MaxLength),
		gType:    GType,
		//EWeight:  make([][]int, MaxLength),
		EWeight:  t,
		IsTravel: make([]int, MaxLength),
	}, nil
}

//  增加一个节点
func (gm *GraphMatrix) AddV(s string) {

	gm.Data[gm.VNum] = s
	gm.VNum++
	return
}

// 增加一条边
func (gm *GraphMatrix) AddE(start, end string, weight int) error {
	// 各种问题
	// 1. 找start end index
	// 2. weight 不能为负数
	// 3. 分邮有向和无向图

	var startIndex, endIndex = -1, -1

	for key, value := range gm.Data {
		if value == start {
			startIndex = key
		} else if value == end {
			endIndex = key
		}
	}

	//for i := 0; i < gm.VNum; i++ {
	//	if gm.Data[i] == start {
	//		startIndex = i
	//	} else if gm.Data[i] == end {
	//		endIndex = i
	//	}
	//}

	if startIndex == -1 || endIndex == -1 {
		return errors.New("start and end parameter invalid")
	}

	if weight <= 0 {
		return errors.New("weight > 0")
	}
	gm.EWeight[startIndex][endIndex] = weight
	gm.ENum++

	if gm.gType == 0 {
		(gm.EWeight)[endIndex][startIndex] = weight
		gm.ENum++
	}

	return nil
}

func (gm *GraphMatrix)depthTravelOne(n int){

	fmt.Println("===> ", gm.Data[n])
	gm.IsTravel[n] = 1

	for i := 0; i < gm.VNum; i++ {
		if gm.EWeight[n][i] != 0 && gm.IsTravel[i] == 0 {
			gm.depthTravelOne(i)
		}
	}
}

//  深度遍历 version 1
func (gm *GraphMatrix) DFS_V1() {
	// 借用遍历标记
	// 初始化IsTravel
	for i := 0; i < gm.VNum; i++{
		gm.IsTravel[i] = 0
	}
	// 遍历每一个

	for i := 0; i < gm.VNum; i++{

		if gm.IsTravel[i] == 0 {
			gm.depthTravelOne(i)
		}
	}
}

// DFS 栈方式
func (gm *GraphMatrix) DFS_Stack() {


}

// 用队列的方式 广度遍历
func (gm *GraphMatrix) BFS_Queue() {

}

//   A B C D
// A
// B
// C
// D
func (gm *GraphMatrix) PrintGraphWeigh() {

	for i := 0; i < gm.VNum; i++ {
		fmt.Printf("\t%v", gm.Data[i])
	}
	fmt.Printf("\n")
	for i := 0; i < gm.VNum; i++ {
		fmt.Printf("%v\t", gm.Data[i])
		for j := 0; j < gm.VNum; j++ {
			fmt.Printf("%v\t", gm.EWeight[i][j])
		}
		fmt.Printf("\n")
	}
}

func CreateGraphMatrie() error {

	var GType = 1 // 选择无向图

	gm, err := NewGraphMatrix(GType)
	if err != nil {
		return errors.New("initial graph failed")
	}

	gm.AddV("A")
	gm.AddV("B")
	gm.AddV("C")
	gm.AddV("D")

	gm.AddE("A", "B", 3)
	gm.AddE("A", "C", 1)
	gm.AddE("B", "C", 2)
	gm.AddE("C", "D", 1)
	gm.AddE("B", "D", 2)
	gm.AddE("A", "D", 1)

	// 输出图的权值表
	gm.PrintGraphWeigh()

	// 从第0个开始遍历
	//var index = 0
	gm.DFS_V1()
	return nil
}

//func main() {
//
//	if err := CreateGraphMatrie(); err != nil {
//		fmt.Println(err.Error())
//	}
//
//}
// 结果
//A       B       C       D
//A       0       3       1       1
//B       3       0       2       2
//C       1       2       0       1
//D       1       2       1       0
//===>  A
//===>  B
//===>  C
//===>  D

