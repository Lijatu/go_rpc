package main

import (
	"context"
	"flag"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()


	d:= client.NewPeer2PeerDiscovery("tcp@"+*addr,"")
	//客户端使用了 Failtry 模式并且随机选择节点。
	//NewXClient 必须使用服务名称作为第一个参数， 然后是 failmode、 selector、 discovery等其他选项。
	xclient := client.NewXClient("Arith",client.Failtry,client.RandomSelect,d,client.DefaultOption)

	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	//特殊情况下，你可以使用 XClient 的 Broadcast 和 Fork 方法。
	//Broadcast 表示向所有服务器发送请求，只有所有服务器正确返回时才会成功。此时FailMode 和 SelectMode的设置是无效的。请设置超时来避免阻塞。
	//Fork 表示向所有服务器发送请求，只要任意一台服务器正确返回就成功。此时FailMode 和 SelectMode的设置是无效的。
	call,err := xclient.Go(context.Background(),"Mul",args,reply,nil)
	if err!=nil{
		log.Fatalf("Failed to call : %v",err)

	}
	 replyCall := <-call.Done

	 if replyCall.Error != nil{
		 log.Fatalf("failed to call: %v", replyCall.Error)
	 }else{
		 log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	 }
}