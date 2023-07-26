package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo"
	"github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo/studentservice"
)

func main() {
	cli, err := studentservice.NewClient("student-server",
		client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic("err init client:" + err.Error())
	}

	resp, err := cli.Register(context.Background(), &demo.Student{
		Id:   1,
		Name: "XueYu",
		College: &demo.College{
			Name:    "SE",
			Address: "NJU",
		},
		Email: []string{
			"211250052@smail.nju.edu.cn",
		},
	})
	if err != nil {
		panic("register failed")
	}

	/*resp, err := cli.Query(context.Background(), &demo.QueryReq{
		Id: 1,
	})

	if err != nil {
		panic("err query:" + err.Error())
	}*/
	klog.Infof("resp: %v", resp)
}
