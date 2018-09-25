package main

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

var frontArrs = []int{1, 2, 4, 7, 3, 5, 6, 8}
var middleArrs = []int{4, 7, 2, 1, 5, 3, 8, 6}


func craeteBinaryTreeNode(fArr, mArr []int, head *BinaryTreeNode){
	// 先序找根
	// 中序找左右

	//fLength := len(fArr)
	//mLength := len(mArr)
	//
	//if fLength

}

func main(){
	//head := craeteBinaryTreeNode(fArr, mArr)
}