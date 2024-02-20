package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class string
	Subjects []string
	// courses as an embedding
}

type Class interface {
	RegisterClass(name string, className string) string
	RegisterSubject(name string, subjectname string) string
	DropClass(courses *[]string, classname string) string
	DropSubject(subjects *[]string, classname string) string
	ViewSubjects(s *Student) []string
}

func (s *Student) RegisterClass(name string, className string) string {
	return fmt.Sprintf("Student: %s, is now registered for %s", name, className)
}

func (s *Student) RegisterSubject(subjectname string) string {
	s.Subjects = append(s.Subjects, subjectname)
	return fmt.Sprintf("Student: %s, is now registered for %s class", s.Name, subjectname)
}

func (s *Student) ViewSubjects() string{
	for _, subject := range s.Subjects{
		fmt.Printf("%v, is taking the following subjects: %v", s.Name, subject)
		return subject
	}
	return ""
}

// func (s *Student) DropClass(courses []string, classname string) *[]map[Student]string {
	
// }

// func (s *Student) DropSubject(subjects *[]map[Student]string, subjectname string) *[]map[Student]string {
// 	// loop through subjects
// 	for i, subject := range *subjects {
// 		// check if subject exists, and then if any subjects match the given subject name
// 		if _, exists := subject[*s]; exists && subject[*s] == subjectname{ // explore using the inbuilt map delete function to remove instead of  the append
// 			*subjects = append((*subjects)[:i],(*subjects)[i+1:]... ) // we want to remove the found subject at element i, that matches the condition

// 		}
// 	}
// 	return subjects
// 	//subject[*s] works because to access the value of the map, we call the key value
	
// }

func main() {
	johnpaul := Student{
		Name: "Johnpaul",
		Age: 20,
		Class: "Medicine",
	}
	// it is not compulsory to include the ampersand as it infers to the pointer
	alexandra := Student{
		Name: "Alexandra",
		Age: 21,
		Class: "PreMed",
	}
	physics:= johnpaul.RegisterSubject("Physics")
	chemistry := alexandra.RegisterSubject("Chemistry")

	biology := johnpaul.RegisterSubject("Biology")

	fmt.Println(johnpaul.ViewSubjects());

	fmt.Println(physics)
	fmt.Println(chemistry)
	fmt.Println(biology)

	fmt.Println(alexandra)
	fmt.Println(johnpaul)

}
// It rained