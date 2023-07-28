package main

import (
	"context"
	"github.com/xueyyyyyyu/rpcsvr/kitex_gen/demo"
	"testing"
)

// Test the Register method of StudentServiceImpl
func TestStudentService_Register(t *testing.T) {
	// Initialize the StudentServiceImpl and test data
	service := &StudentServiceImpl{}
	student := &demo.Student{
		Id:      1,
		Name:    "John Doe",
		Sex:     "Male",
		Age:     25,
		College: &demo.College{Name: "SE", Address: "NJU"},
		Email:   []string{"john.doe@example.com"},
	}

	len1 := len(id2Student)

	// Test registering a new student
	_, err := service.Register(context.Background(), student)
	if err != nil {
		t.Errorf("Failed to register student: %v", err)
	}

	// Verify the student is present in the map
	if len(id2Student) != 1+len1 {
		t.Errorf("Expected a new student in the map, but found %d", len(id2Student)-len1)
	}
}

// Test the Query method of StudentServiceImpl
func TestStudentService_Query(t *testing.T) {
	// Initialize the StudentServiceImpl and test data
	service := &StudentServiceImpl{}
	student := &demo.Student{
		Id:      1,
		Name:    "John Doe",
		Sex:     "Male",
		Age:     25,
		College: &demo.College{Name: "SE", Address: "NJU"},
		Email:   []string{"john.doe@example.com"},
	}
	id2Student[1] = *student

	// Test querying an existing student
	req := &demo.QueryReq{Id: 1}
	resp, err := service.Query(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to query student: %v", err)
	}

	// Verify the response is correct
	if resp == nil {
		t.Errorf("Expected student response, but got nil")
	} else {
		expectedName := "John Doe"
		if resp.Name != expectedName {
			t.Errorf("Expected student name %q, but got %q", expectedName, resp.Name)
		}
	}

	// Test querying a non-existing student
	req = &demo.QueryReq{Id: 2}
	resp, err = service.Query(context.Background(), req)
	if err == nil || resp != nil {
		t.Errorf("Expected 'not found' error and nil response, but got: %v, %v", err, resp)
	}
}
