package htest

import (
	"io"
	"net/http"
	"fmt"
)

// 这是对hello 的处理
func HandleHello(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	io.WriteString(res, "你好")
}
func handleHello2(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	io.WriteString(res, "你好")
}
