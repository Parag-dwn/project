package storage

import (
	"github.com/Parag-dwn/student-api/internal/types"
)

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	PutStudentById(name string, email string, age int, id int64) (int64, error)
	DeleteStudentById(id int64) error
}
