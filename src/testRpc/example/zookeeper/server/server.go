package main

import (
	"flag"
	example "github.com/rpcxio/rpcx-examples"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"

)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/youpin/services", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegisterPlugins(s)

	s.RegisterName("Arith",new(example.Arith),"")
	s.Serve("tcp",*addr)
}

func addRegisterPlugins(s *server.Server){
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress: "tcp@"+*addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath: *basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil{
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}