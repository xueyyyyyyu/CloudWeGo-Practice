package main

import (
	"context"
	demo "github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// map 缓存
var id2Student = map[int]demo.Student{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	id := int(student.Id)
	_, found := id2Student[id]
	if found {
		return
	} else {
		id2Student[id] = *student
	}
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	/*resp = &demo.Student{
		Id:   1,
		Name: "XueYu",
		College: &demo.College{
			Name:    "SE",
			Address: "NJU",
		},
		Email: []string{
			"211250052@smail.nju.edu.cn",
		},
	}*/

	student, found := id2Student[int(req.Id)]
	if found {
		resp = &student
	} else {
		resp = &demo.Student{
			Name: "not found",
		}
	}

	return
}
