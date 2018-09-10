
package interview

import (
	"testing"
	"fmt"
	"reflect"
)

func TestGetOrderInt(t *testing.T){
	var m = map[string]int{
		"c":1,
		"b":2,
		"a":3,
	}

	res := GetOrderInt(m)
	fmt.Println("res", res)
	if reflect.DeepEqual(res, []int{3,2,1}){
		t.Log("校验通过")
	}else{
		t.Fail()
	}
}