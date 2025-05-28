package test

import (
	"github.com/annuums/go-study-web-server/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StudentServiceSuite struct {
	suite.Suite
	svc *services.StudentServiceImpl
}

func (s *StudentServiceSuite) SetupTest() {
	// 매 테스트 전에 실행됨
	s.svc = services.NewStudentService().(*services.StudentServiceImpl)
	_ = s.svc.AddStudent(&services.Student{ID: 1, Name: "Alice", Grade: 1})
}

func (s *StudentServiceSuite) TestAddAndGet() {
	stu := &services.Student{ID: 10, Name: "Carol", Grade: 1}
	err := s.svc.AddStudent(stu)
	s.Require().NoError(err)

	got, err := s.svc.GetStudent(10)
	s.Require().NoError(err)
	assert.Equal(s.T(), stu.Name, got.Name)
}

func (s *StudentServiceSuite) TestDelete() {
	stu := &services.Student{ID: 20, Name: "Dave", Grade: 4}
	_ = s.svc.AddStudent(stu)

	err := s.svc.DeleteStudent(20)
	s.Require().NoError(err)
	_, err = s.svc.GetStudent(20)
	assert.Error(s.T(), err)
}

func (s *StudentServiceSuite) TestDuplicate() {
	stu := &services.Student{ID: 1, Name: "Alice", Grade: 1}
	err := s.svc.AddStudent(stu)

	s.Require().Error(err)
	assert.Error(s.T(), err)
}

func TestStudentServiceSuite(t *testing.T) {
	suite.Run(t, new(StudentServiceSuite))
}
