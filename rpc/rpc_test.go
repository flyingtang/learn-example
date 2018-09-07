package server

import (
	"net/rpc"
	"testing"
)

func TestNewRpcServer(t *testing.T) {
	NewRpcServer()
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:5000")
	if err != nil {
		t.Error(err.Error())
	}
	var (
		args = &Args{5, 5}
		rep  int
	)

	err = client.Call("Arith.Multiply", args, &rep)
	if err != nil {
		t.Error(err.Error())
	}

	if rep == 25 {
		t.Log("校验通过")
	} else {
		t.Error("校验失败")
	}
}
