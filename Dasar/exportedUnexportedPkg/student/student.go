package student

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Name init"
	Student.Grade = 100

	fmt.Println("--> student/student.go imported")
}
