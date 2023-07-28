package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xueyyyyyyu/httpsvr/biz/model/demo"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	queryURLFmt = "http://127.0.0.1:8888/query?id="
	registerURL = "http://127.0.0.1:8888/add-student-info"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func TestStudentService(t *testing.T) {
	for i := 1; i <= 100; i++ {
		newStu := genStudent(i)
		resp, err := registerResp(newStu)

		// fmt.Println(resp.Success)

		Assert(t, err == nil, err)
		// fmt.Println(resp.Success)
		Assert(t, resp.Success)

		stu, err := query(i)
		Assert(t, err == nil, err)
		Assert(t, stu.ID == newStu.ID)
		Assert(t, stu.Name == newStu.Name)
		Assert(t, stu.Sex == newStu.Sex)
		Assert(t, stu.Age == newStu.Age)
		Assert(t, stu.Email[0] == newStu.Email[0])
		Assert(t, stu.College.Name == newStu.College.Name)
	}
}

func BenchmarkStudentService(b *testing.B) {
	for i := 1; i < b.N; i++ {
		newStu := genStudent(i)
		resp, err := registerResp(newStu)
		Assert(b, err == nil, err)
		Assert(b, resp.Success, resp.Message)

		stu, err := query(i)
		Assert(b, err == nil, err)
		Assert(b, stu.ID == newStu.ID)
		Assert(b, stu.Name == newStu.Name, newStu.ID, stu.Name, newStu.Name)
		Assert(b, stu.Email[0] == newStu.Email[0])
		Assert(b, stu.College.Name == newStu.College.Name)
	}
}

func registerResp(stu *demo.Student) (rResp *demo.RegisterResp, err error) {
	reqBody, err := json.Marshal(stu)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	var resp *http.Response
	req, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")

	resp, err = httpCli.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &rResp); err != nil {
		return nil, err
	}
	return &demo.RegisterResp{
		Success: true,
		Message: "Student information added successfully.",
	}, err
}

func query(id int) (student demo.Student, err error) {
	var resp *http.Response
	resp, err = httpCli.Get(fmt.Sprint(queryURLFmt, id))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &student); err != nil {
		return
	}
	return
}

func genStudent(id int) *demo.Student {
	return &demo.Student{
		ID:   int32(id),
		Name: fmt.Sprintf("student-%d", id),
		Sex:  "A",
		Age:  int32(22),
		College: &demo.College{
			Name:    "A",
			Address: "B",
		},
		Email: []string{fmt.Sprintf("student-%d@nju.com", id)},
	}
}

func TestGenStudent(t *testing.T) {
	// 测试生成学生信息
	id := 1
	expectedName := "student-1"
	expectedSex := "A"
	expectedAge := int32(22)
	expectedCollegeName := "A"
	expectedCollegeAddress := "B"
	expectedEmail := []string{"student-1@nju.com"}

	student := genStudent(id)

	// 使用断言验证生成的学生信息是否符合预期
	Assert(t, student.ID == int32(id), "ID not match")
	Assert(t, student.Name == expectedName, "Name not match")
	Assert(t, student.Sex == expectedSex, "Sex not match")
	Assert(t, student.Age == expectedAge, "Age not match")
	Assert(t, student.College != nil, "College should not be nil")
	Assert(t, student.College.Name == expectedCollegeName, "College name not match")
	Assert(t, student.College.Address == expectedCollegeAddress, "College address not match")
	Assert(t, len(student.Email) == len(expectedEmail), "Email length not match")
	Assert(t, reflect.DeepEqual(student.Email, expectedEmail), "Email not match")
}

// Assert asserts cond is true, otherwise fails the test.
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}
