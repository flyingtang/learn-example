package server

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divize by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func NewRpcServer() {
	rpc.Register(new(Arith))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		panic(err.Error())
		return
	}
	go http.Serve(l, nil)
}
