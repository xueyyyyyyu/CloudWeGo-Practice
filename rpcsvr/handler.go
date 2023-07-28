package main

import (
	"context"
	"errors"
	demo "github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo"
	"sync"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	mu sync.RWMutex // Define a mutex to protect id2Student
}

// map 缓存
var id2Student = map[int]demo.Student{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	s.mu.Lock()         // Lock the mutex before modifying the map
	defer s.mu.Unlock() // Unlock the mutex after the function exits

	// rpc register
	id := int(student.Id)
	_, found := id2Student[id]
	if found {
		return
	} else {
		id2Student[id] = *student
	}
	// fmt.Println(id2Student)
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	s.mu.Lock()         // Lock the mutex before modifying the map
	defer s.mu.Unlock() // Unlock the mutex after the function exits

	// rpc query
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

	// fmt.Println(id2Student)

	student, found := id2Student[int(req.Id)]
	if found {
		resp = &student
	} else {
		err = errors.New("not found")
	}

	// fmt.Println(resp)

	return
}
