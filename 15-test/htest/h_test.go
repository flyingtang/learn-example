package htest

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"io/ioutil"
)

func TestHandleHello(t *testing.T) {
	//httptest.NewServer()

	tests := []struct {
		h http.HandlerFunc
		code int
		message string
	}{
		{HandleHello,
		200,
		"123"},}

		for _, tt := range tests{
			ress := httptest.NewRecorder()
			reqs := httptest.NewRequest(http.MethodGet,"http://ixiaotang.cn", nil)
			tt.h(ress,reqs)
			b, _ := ioutil.ReadAll(ress.Body)
			body := string(b)
			fmt.Println(body, ress.Code)
		}
}

func ExampleHandleHello() {
 // 例子
 // Output:
 // eg :1...
}