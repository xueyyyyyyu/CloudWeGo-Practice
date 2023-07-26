// Code generated by hertz generator.

package demo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xueyyyyyyu/httpsvr/biz/model/demo"
	"log"
)

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//step 3.1: 非泛化调用的register
	/*cli, err := studentservice.NewClient("student-server",
		kclient.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic("err init client" + err.Error())
	}

	resp, err := cli.Register(context.Background(),
		&kdemo.Student{
			Id:   req.ID,
			Name: req.Name,
			College: &kdemo.College{
				Name:    req.College.Name,
				Address: req.College.Address,
			},
			Email: req.Email,
		})
	c.JSON(consts.StatusOK, resp)*/

	// step 3.2: 泛化调用的register
	cli := initGenericClient()
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("get custom req failed")
	}
	resp, err := cli.GenericCall(ctx, "Register", customReq)
	if err != nil {
		panic("generic call failed")
	}
	c.JSON(consts.StatusOK, resp)
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.QueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	/*cli, err := studentservice.NewClient("student-server",
		kclient.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic("err init client:" + err.Error())
	}

	resp, err := cli.Query(context.Background(), &kdemo.QueryReq{
		Id: 1,
	})
	if err != nil {
		panic("err query:" + err.Error())
	}

	c.JSON(consts.StatusOK, resp)*/

	// 泛化调用的query
	cli := initGenericClient()
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("get custom req failed")
	}
	resp, err := cli.GenericCall(ctx, "Query", customReq)
	if err != nil {
		panic("generic call failed")
	}
	realResp := resp.(*generic.HTTPResponse)
	c.JSON(consts.StatusOK, realResp.Body)
}

// 泛化调用
func initGenericClient() genericclient.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	p, err := generic.NewThriftFileProvider("../student.thrift")
	if err != nil {
		panic(err)
	}

	// 构造HTTP类型的泛化调用
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	cli, err := genericclient.NewClient("destServiceName", g,
		kclient.WithHostPorts("127.0.0.1:8889"),
		kclient.WithResolver(r))
	if err != nil {
		panic(err)
	}

	return cli
}
