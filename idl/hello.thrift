// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

struct OtherReq {
    1: string Other (api.body="other");
}

struct OtherResp {
    1: string Resp;
}

/* TODO task2
使用 IDL，定义一个接口，通过 hz 工具基于 IDL 生成代码:
要求:
1. 定义一个 GET 接口用于查询学生信息，和一个 POST 接口用于写入学生信息
2. 学生信息由学号（ID）、名字（Name）、喜欢的食物（Favorite）三个字段构成
2. GET 接口的参数使用 api.query 注解从 query 字段绑定对应的 ID 值
3. POST 接口使用 api.json 注解从 JSON 格式的 body 中绑定对应的学生信息

实现效果举例：
1. POST 请求 url：127.0.0.1:8888/add-student-info, body 为 {“ID”:1,”Name”: “Sam”, “Favorite”: “apple”}的 JSON 数据，将ID为1的学生信息记录到服务端
2. GET 请求 url：127.0.0.1:8888/query?id=1，返回id为1的上述学生信息
*/


struct AddStudentInfoReq {
    1: i32 ID (api.body="ID");
    2: string name (api.body="Name");
    3: string favourite (api.body="Facourite");
}

struct AddStudentInfoResp {
    1: string resp;
}

struct QueryReq{
    1: i32 ID (api.query="id");
}

struct QueryResp {
    1: string resp (api.body="res");
}



service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    OtherResp OtherMethod(1: OtherReq request) (api.post="/other");
}

service NewService {
    HelloResp NewMethod(1: HelloReq request) (api.get="/new");
}

service StudentServive {
    QueryResp QueryMethod(1: QueryReq request) (api.get="/query");
    AddStudentInfoResp WriteMethod(1: AddStudentInfoReq request) (api.post="/add-student-info");
}