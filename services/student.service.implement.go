package services

import (
	"fmt"
	"sync"
)

type StudentServiceImpl struct {

	// students 메모리에서 학생 정보를 저장하는 맵입니다.
	// Key: 학번, Value: Student
	students *map[int]Student

	// mu 동시 접근을 위한 읽기/쓰기 잠금입니다.
	// Golang은 map이 기본적으로 Thread Safe하지 않기 때문에 설정해요.
	mu sync.RWMutex
}

func NewStudentService() StudentService {

	students := make(map[int]Student)

	return &StudentServiceImpl{
		students: &students,
	}
}

// GetStudent 메모리에서 ID에 해당하는 학생 정보를 가져옵니다.
func (s *StudentServiceImpl) GetStudent(id int) (*Student, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	// 여기에 학생 정보를 가져오는 로직을 구현합니다.
	// 예시로, ID가 1인 학생 정보를 반환하는 로직은 생략합니다.
	if std, ok := (*s.students)[id]; ok {
		return &std, nil
	}

	return nil, fmt.Errorf("학생을 찾을 수 없습니다 :: %d", id)
}

// AddStudent 메모리에 새로운 학생 정보를 추가합니다.
func (s *StudentServiceImpl) AddStudent(student *Student) error {

	s.mu.RLock()
	if _, ok := (*s.students)[student.ID]; ok {
		s.mu.RUnlock()

		return fmt.Errorf("이미 존재하는 학생입니다 :: %d", student.ID)
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	(*s.students)[student.ID] = *student
	fmt.Printf("학생 정보 추가됨 :: ID: %d, Name: %s, Grade: %d\n", student.ID, student.Name, student.Grade)

	return nil // 성공적으로 추가되었다고 가정
}

// DeleteStudent 메모리에서 ID에 해당하는 학생 정보를 삭제합니다.
func (s *StudentServiceImpl) DeleteStudent(id int) error {

	// 학생 정보를 삭제하기 전에 해당 ID가 존재하는지 확인합니다.
	s.mu.RLock()
	if _, ok := (*s.students)[id]; !ok {

		s.mu.RUnlock()
		return fmt.Errorf("삭제할 학생을 찾을 수 없습니다 :: %d", id)
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	// 학생 정보를 삭제합니다.
	delete(*s.students, id)
	fmt.Printf("학생 정보 삭제됨 :: ID: %d\n", id)

	return nil // 해당 ID의 학생이 없을 경우 nil 반환
}
