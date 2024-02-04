package students

import (
	"errors"
)

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"Name"`
	IDNumber string `json:"Number"`
	School   string `json:"School"`
}

var studentList = []Student{
	{ID: "1", Name: "Dauren", IDNumber: "22B031142", School: "SITE"},
	{ID: "2", Name: "Bob", IDNumber: "21B031152", School: "BS"},
	{ID: "3", Name: "Charlie", IDNumber: "23B031146", School: "FEOGI"},
	{ID: "4", Name: "Diana", IDNumber: "20B031145", School: "MCM"},
	{ID: "5", Name: "Edward", IDNumber: "21B031141", School: "ISE"},
}

func health_check() string {
	return "This is the Students package"
}

func GetAllStudents() []Student {
	return studentList
}

func GetStudentByID(id string) (*Student, error) {
	for _, student := range studentList {
		if student.ID == id {
			return &student, nil
		}
	}
	return nil, errors.New("student not found")
}