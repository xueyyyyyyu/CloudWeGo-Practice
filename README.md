# 说明文档

> 代码仓库：[xueyyyyyyu/CloudWeGo-Practice](https://github.com/xueyyyyyyu/CloudWeGo-Practice)
>
> 队伍名称：终于做队
>
> 队伍成员：薛宇 211250052

## 1. 环境配置

VMware Workstation Pro 17 + Ubuntu 20.04 + Goland

### 1.1 golang安装

> go version go1.20.5 linux/amd64

1. 官网下载 1.20.5 版本的压缩包

2. 解压缩 `sudo tar -C /usr/local -zxvf go1.20.5.linux-amd64.tar.gz`

3. 配置代理：执行命令 `go env -w GOPROXY=http://goproxy.cn`

4. 执行命令 `go version` 查看到版本信息即为安装成功

5. 将$GOPATH/bin添加到PATH环境变量（go install 安装的命令会在$GOPATH/bin目录下）

   ```shell
   sudo vim ~/.bashrc
   #添加以下配置并保存退出
   export PATH=$GOPATH/bin:$PATH
   # 激活配置
   source ~/.bashrc
   ```

### 1.2 hertz安装

> hz version v0.6.5

1. 执行命令 `go install -v github.com/cloudwego/hertz/cmd/hz@latest`
2. `hz --version` 查看到版本信息即为安装成功

### 1.3 thriftgo安装

> thriftgo 0.2.12

1. 执行命令`GO111MODULE=on go install github.com/cloudwego/thriftgo@latest`
1. `thriftgo version` 查看版本

### 1.4 kitex安装

> v0.6.1

1. 执行命令 ` go install github.com/cloudwego/kitex/tool/cmd/kitex@latest`
2. `kitex --version` 查看版本

### 1.5 etcd安装

> etcd Version: 3.5.9
> Git SHA: bdbbde998
> Go Version: go1.19.9
> Go OS/Arch: linux/amd64
>
> etcdctl version: 3.5.9
> API version: 3.5

1. 下载：https://github.com/etcd-io/etcd/releases
2. 解压得到 etcd 和 etcdctl，将这两个文件复制到 usr/local/bin
3. `etcd --version` , `etcdctl version` 查看版本

### 1.6 Apache Benchmark安装（性能测试）

1. `sudo apt install apache2-utils`

## 2. API网关核心功能

### 2.1 部署步骤

1. rpcsvr 目录下运行`sh build.sh` 和 `sh output/bootstrap.sh` 运行 RPC 服务端
2. 新建终端，httpsvr 目录下运行 `go build` 和 `./httpsvr` 运行 HTTP 客户端
3. 新建终端，任意目录下运行 `etcd --log-level debug` 运行 etcd
4. 以上终端都不要关闭，需要关闭时使用`ctrl + c`
5. 新建终端进行POST或者GET操作



### 2.2 接口文档

#### 2.2.1 注册学生信息接口

请求URL：/add-student-info

请求方法：POST

请求参数：

| 参数名  | 类型     | 必需 | 示例值                                         | 描述     |
| ------- | -------- | ---- | ---------------------------------------------- | -------- |
| id      | int32    | 是   | 100                                            | 学生ID   |
| name    | string   | 是   | "XueYu"                                        | 学生姓名 |
| sex     | string   | 是   | "male"                                         | 学生性别 |
| age     | int32    | 是   | 20                                             | 学生年龄 |
| college | *collage | 是   | {"name": "software college", "address": "NJU"} | 所在院系 |
| email   | []string | 否   | ["211250052@smail.nju.edu.cn"]                 | 学生邮箱 |

请求示例：

```go
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info -d '{"id": 100, "name":"XueYu", "sex":"male", "age":20, "college": {"name": "software college", "address": "NJU"}, "email": ["211250052@smail.nju.edu.cn"]}'
```

响应结果：

- 如果 id 已经存在则不插入

  ```GO
  {"Header":{},"StatusCode":0,"Body":{"message":"Student ID already exists.","success":false},"GeneralBody":null,"ContentType":"application/json","Renderer":{}}
  ```

- 否则插入

  ```GO
  {"Header":{},"StatusCode":0,"Body":{"message":"Student information added successfully.","success":true},"GeneralBody":null,"ContentType":"application/json","Renderer":{}}
  ```

  

#### 2.2.2 查询学生信息接口

请求URL：/query?id={id}

请求方法：GET

请求参数：

| 参数名 | 类型  | 必需 | 示例值 | 描述   |
| ------ | ----- | ---- | ------ | ------ |
| id     | int32 | 是   | 100    | 学生ID |

请求示例：

```go
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/query?id=100
```

响应数据：

```GO
{"age":20,"college":{"address":"NJU","name":"software college"},"email":["211250052@smail.nju.edu.cn"],"id":100,"name":"XueYu","sex":"male"}
```



### 2.3 HTTP POST请求的正确接收与响应

```Go
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info -d '{"id": 100, "name":"XueYu", "sex":"male", "age":20, "college": {"name": "software college", "address": "NJU"}, "email": ["211250052@smail.nju.edu.cn"]}'
```

终端运行上面的命令后在浏览器打开http://127.0.0.1:8888/query?id=100或者在终端中运行：

```Go
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/query?id=100
```

可以得到POST的正确结果



### 2.4 根据请求路由确认目标服务与方法

> 服务粒度的流量路由

在前面HTTP POST的验证中，/add-student-info和/query路由分别对应studentservice的Register和Query方法



### 2.5 网关内的IDL管理模块

> 热更新：不重启 httpsvr 就能体现 idl 更新的影响

修改 student.thrift 文件，添加 student 的 age 字段，在 rpcsvr 目录下运行`kitex -module github.com/xueyyyyyyu/rpcsvr -service student-server ../student.thrift` 更新 kitex server 后不重启 httpsvr，终端输入：

```GO
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info -d '{"id": 100, "name":"XueYu", "sex":"male", "age":20, "college": {"name": "software college", "address": "NJU"}, "email": ["211250052@smail.nju.edu.cn"]}'
```

浏览器访问 http://127.0.0.1:8888/query?id=100 得到了含有 age 的结果，说明热更新成功实现。



### 2.6 Kitex泛化调用客户端

> 泛化：没有生成代码，没有业务数据结构。基于idl中的注解映射。
>

代码中 httpsvr/biz/handler/demo/student_service.go 的 initGenericClient 方法实现了远程客户端的创建，并且注册到 etcd 中，采用加权随机的负载均衡。在 Register 方法和 Query 方法中调用 initGenericClient 生成 kclient（即Kitex的client），通过httpReq、customReq发起请求，cli.GenericCall(ctx, "方法名", customReq) 调用指定的方法，实现了Kitex泛化调用客户端。



### 2.7 编码：可读性、模块划分、单测覆盖

可读性：关键代码有注释，编码风格还行

模块划分：基本按照下图![整体设计](.\README\image-20230728002719591.png)

单测覆盖：

- rpcsvr 中 handler_test.go 对 handler.go 中的 Register 和 Query 方法进行测试
- httpsvr 中 main_test.go 对 studentService 进行测试
- httpsvr/main_test.go 中 TestGenStudent 方法对 genStudent 方法进行测试



## 3. 性能测试和优化报告

### 3.1 测试方案说明

- ab 测试：虚拟机是四核，所以开启四线程。请求数定为 1000。在 etcd，rpcsvr，httpsvr 成功运行的情况下，先进行一次 registerStudent 操作，然后运行：
  - `ab -n 1000 -c 4 http://127.0.0.1:8888/ping` 对路由 /ping 进行测试
  - `ab -n 1000 -c 4 http://127.0.0.1:8888/query?id=100` 对路由 /query 进行测试

- golang 性能测试：测试 register 和 query 方法
  - 串行测试：在 httpsvr/main_test.go 中运行 BenchmarkStudentService 方法
  - 并行测试：在 httpsvr/main_test.go 中运行 BenchmarkStudentServiceParallel 方法
  - 观测内存：在 httpsvr 目录下运行 `go test -bench=. -benchmem main_test.go `



### 3.2 性能测试数据

> 详细测试数据见 testData 目录，多次进行相同测试结果并不一样

- 对于 1000 个 ping 请求，耗时 2 ms 以内完成，结果如 ping.txt 所示

- 对于 1000 个 query 请求，耗时 30 ms 以内完成，结果如 query.txt 所示

- 串行测试结果：

  ```
  goos: linux
  goarch: amd64
  pkg: github.com/xueyyyyyyu/httpsvr
  cpu: 11th Gen Intel(R) Core(TM) i5-1155G7 @ 2.50GHz
  BenchmarkStudentService
  BenchmarkStudentService-4   	     308	   5766529 ns/op
  PASS
  ```

- 并行测试结果：

  ```
  goos: linux
  goarch: amd64
  pkg: github.com/xueyyyyyyu/httpsvr
  cpu: 11th Gen Intel(R) Core(TM) i5-1155G7 @ 2.50GHz
  BenchmarkStudentServiceParallel
  BenchmarkStudentServiceParallel-4   	     866	   3072547 ns/op
  PASS
  ```

- 内存测试结果：

  ```
  goos: linux
  goarch: amd64
  cpu: 11th Gen Intel(R) Core(TM) i5-1155G7 @ 2.50GHz
  BenchmarkStudentService-4                    385           3203811 ns/op           10645 B/op          152 allocs/op
  BenchmarkStudentServiceParallel-4            559           2619191 ns/op           10814 B/op          151 allocs/op
  PASS
  ok      command-line-arguments  3.598s
  
  ```

  

### 3.3 优化方案说明：未完成

### 3.4 优化后性能数据：未完成

## 4. 附加要求：未完成

## 5. 作业要求

![基础要求](.\README\image-20230724224511011.png)

![附加要求](.\README\image-20230724224617096.png)

