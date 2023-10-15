package models

import (
	"fmt"
	"go_algo/db/connect"
	"log"

	"github.com/google/uuid"
)

var db, _ = connect.GetDB()

/*
@typescript type:

	type StudentModel = {
		id: string,
		student_name: string,
		email_id: string,
		student_address?: string
	};
*/
type Student struct {
	UiD     string  `gorm:"primaryKey;<-:create;not null;unique;column:student_id" json:"id"`
	Name    string  `gorm:"not null;" json:"student_name"`
	Email   string  `gorm:"not null;unique_index;" json:"email_id"`
	Address *string `gorm:"size:255" json:"student_address,omitempty"`
}

func FetchAllStudents() []Student {
	students := make(chan []Student)

	go func(channel chan<- []Student) {
		studentList := []Student{}

		if result := db.Find(&studentList); result.Error != nil {
			log.Fatal("unable to retrieve student records", result.Error)
		} else {
			channel <- studentList
		}
	}(students)

	return <-students
}

func CreateNewStudent(name string, email string, address *string) {

	returnValue := db.Create(Student{UiD: uuid.New().String(), Name: name, Email: email, Address: address})

	fmt.Println(returnValue)
}
