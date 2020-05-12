package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"

	example "github.com/rpcxio/rpcx-examples"
)
var (
	addr = flag.String("addr", "localhost:8972", "server address")
)
type Arith struct {
}

func (t *Arith) Mul(ctx context.Context,args example.Args,reply *example.Reply) error{
	reply.C = args.A*args.B
	fmt.Println("C=",reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()

	s.RegisterName("Arith",new(Arith),"")
	err := s.Serve("tcp",*addr)
	if err!=nil{
		panic(err)
	}

}