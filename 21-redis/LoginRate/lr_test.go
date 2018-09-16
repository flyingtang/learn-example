package LoginRate

import (
	"testing"
	"time"
)

func TestIsBeyondMax(t *testing.T) {

	c := NewRedis(":6379")
	defer c.Close()
	testDatas := map[string]interface{}{
		"lastLoginUnix": time.Now().Unix(),
		"count":         7,
	}
	id := "1"
	cs := c.HMSet("1", testDatas)
	if cs.Val() == "OK" {
		t.Log("插入初始数据成功")
	}

	is , err := IsBeyondMax(id, c)
	if err != nil {
		t.Error("判断出错")
	}else{
		if is {
			t.Log("已经超越限制", is)
		} else {
			t.Log("没有超越限制", is)
		}

	}
}
