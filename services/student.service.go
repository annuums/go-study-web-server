package services

type StudentService interface {

	// GetStudent 메모리에서 ID에 해당하는 학생 정보를 가져옵니다.
	GetStudent(id int) (*Student, error)

	// AddStudent 메모리에 새로운 학생 정보를 추가합니다.
	AddStudent(student *Student) error

	// DeleteStudent 메모리에서 ID에 해당하는 학생 정보를 삭제합니다.
	DeleteStudent(id int) error
}
