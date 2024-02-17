package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class string
	// courses as an embedding
}

type Class interface {
	RegisterClass(name string, className string) string
	RegisterSubject(name string, className string) string
	DropClass(courses *[]map[Student]string, classname string) string
	DropSubject(subjects *[]map[Student]string, classname string) string
}


type Courses struct {

}

// I could write student methods using pointers.
func (s *Student) RegisterClass(name string, className string) string {
	return fmt.Sprintf("Student: %s, is now registered for %s", name, className)
}

func (s *Student) RegisterSubject(name string, className string) string {
	return fmt.Sprintf("Student: %s, is now registered for %s class", name, className)
}

func (s *Student) DropClass(courses *[]map[Student]string, classname string) {
	for i, course := range *courses {
		if _, exists := course[*s]; exists && course[*s] == classname {// first checks if the course exists, then checks if it exists and if it is equal to the classname to be dropped
			*courses = append((*courses)[:i], (*courses)[i+1:]...) // course[*s] == classname ensures that the subject associated with the student matches the provided classname
			return // and then the append function, takes all the elements up until the present element and then everything after it, and appends skipping the present element.
		}
	}
}

func (s *Student) DropSubject(subjects *[]map[Student]string, subjectname string) {
	for i, course := range *subjects {
		if _, exists := course[*s]; exists && course[*s] == subjectname {
			*courses = append() // it smells like rain, so I'm going to bed now
		}
	}
}

func main() {
	newStudent := &Student{
		Name: "Johnpaul",
		Age: 20,
		Class: "A",
	}

	courses := []map[Student]string{
		{Student{Name: "John", Age: 20, Class: "A"}: "Math"},
		{Student{Name: "John", Age: 20, Class: "A"}: "English"},
	}
}

