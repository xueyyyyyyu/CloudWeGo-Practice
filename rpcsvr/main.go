package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	demo "github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo/studentservice"
	"log"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := demo.NewServer(new(StudentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "student"}), /*,
		server.WithRegistryInfo(&registry.Info{
			Tags: map[string]string{
				"Cluster": "StudentCluster",
			}}), server.WithExitWaitTime(time.Minute)*/)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
