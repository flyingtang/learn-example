package function

import (
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {
	a1 := a()
	res := 1
	fmt.Println("a1 = ", a1)
	if a1 == res {
		t.Log("成功")
	} else {
		t.Fatal("失败")
	}
	b1 := b()
	res = 5
	fmt.Println("b1 = ", b1)
	if b1 == res {
		t.Log("成功")
	} else {
		t.Fatal("失败")
	}
	c1 := c()
	res = 1
	fmt.Println("c1 = ", c1)
	if c1 == res {
		t.Log("成功")
	} else {
		t.Fatal("失败")
	}
}
