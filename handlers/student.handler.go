package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/annuums/go-study-web-server/services"
	"net/http"
	"strconv"
)

type StudentHandler struct {

	// StudentService
	service services.StudentService
}

func (handler *StudentHandler) Handles(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handler.getStudent(res, req)
	case http.MethodPost:
		handler.addStudent(res, req)
	case http.MethodDelete:
		handler.deleteStudent(res, req)
	}
}

func (handler *StudentHandler) getStudent(res http.ResponseWriter, req *http.Request) {

	var studentId int
	var err error
	studentIdStr := req.URL.Query().Get("id")
	if studentId, err = strconv.Atoi(studentIdStr); err != nil {

		http.Error(res, "잘못된 학생 번호 입니다.", http.StatusBadRequest)
		return
	}

	student, err := handler.service.GetStudent(studentId)
	if err != nil {

		fmt.Printf("서버에 오류가 발생했습니다. %d: %v\n", studentId, err)
		http.Error(res, "서버에 오류가 발생했습니다.", http.StatusInternalServerError)
		return
	}
	//fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home")
	fmt.Fprintf(res, "학생 정보: ID: %d, Name: %s, Grade: %d\n", student.ID, student.Name, student.Grade)
}

func (handler *StudentHandler) addStudent(res http.ResponseWriter, req *http.Request) {

	var student services.Student
	var err error

	if err := json.NewDecoder(req.Body).Decode(&student); err != nil {
		http.Error(res, "잘못된 요청 바디: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	if err = handler.service.AddStudent(&student); err != nil {

		http.Error(res, "학생 추가 실패: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "학생 정보 추가됨 :: ID: %d, Name: %s, Grade: %d\n", student.ID, student.Name, student.Grade)
}

func (handler *StudentHandler) deleteStudent(res http.ResponseWriter, req *http.Request) {

	var studentId int
	var err error
	studentIdStr := req.URL.Query().Get("id")
	if studentId, err = strconv.Atoi(studentIdStr); err != nil {

		http.Error(res, "잘못된 학생 번호 입니다.", http.StatusBadRequest)
		return
	}

	if err = handler.service.DeleteStudent(studentId); err != nil {

		http.Error(res, "학생 삭제 실패: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "학생 정보 삭제됨 :: ID: %d\n", studentId)
}
