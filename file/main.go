package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if f, err := os.Open("file/file.txt"); err != nil {
		panic(err.Error())
	} else {
		// 关闭
		defer func() {

			if err := f.Close(); err != nil {
				panic("关闭文件失败")
			}
			log.Fatalln("正常关闭文件")
		}()

		buf := make([]byte, 3)
		for {

			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					log.Println("文件读取完毕")
				} else {
					log.Fatal("文件读取异常")
				}
				break
			} else {
				fmt.Println("内容 ： ", string(buf[:n]))
			}
		}

	}

}
