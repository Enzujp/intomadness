package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   string
	Class string
	// courses as an embedding
}

type Class interface {
	RegisterClass(name string, className string) string
	RegisterSubject(name string, className string) string
	DropClass(courses *[]map[Student]string, classname string) string
	DropSubject(name string, classname string) string
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
		if _, exists := course[*s]; exists && course[*s] == classname {
			*courses = append((*courses)[:i], (*courses)[i+1:]...)
			return
		}
	}
}

func (s *Student) DropSubject(subjects *[]map[Student]string, sunjectname string) {

}
