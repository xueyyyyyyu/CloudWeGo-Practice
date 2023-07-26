package main

import (
	"github.com/cloudwego/kitex/server"
	demo "github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo/studentservice"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
