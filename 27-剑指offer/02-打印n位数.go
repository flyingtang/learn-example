package main

// 打印1 - n 位最大输的十进制


func print1ToMaxOfNdigits(n int){


	// 负值问题
	if n <=0 {
		return
	}


	// 功能问题
	// 初始化 n 位数
	var min, max string
	for i:=0; i < n; i++{
		min +="0"
		max +="9"
	}

	temp := min
	// 边界问题 边界不确定，所有不嫩够用具体的类型来
	// 利用 字符串 或者 数组 来解决
	//1、 有没有达到最大数
	//2、 如何打印

	for noBeyondMax(&temp) {
		printTemp(temp)
	}

}

func noBeyondMax(temp *string) bool{


	// 如果最低位产生了进位，那就说明到了最大啦
	return false
}

func printTemp(temp string) {

	// 有0 要去掉

}