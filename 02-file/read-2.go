package main 

import (
	"bufio"
	"fmt"
	"os"
)

// 读取一行一行的数据

const filePath = "file.txt"


func readLine(){
	f, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(" ===> ", s.Text())
	}
}
