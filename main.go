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
// try to populate this
}

// I could write student methods using pointers.
func (s *Student) RegisterClass(name string, className string) string {
	return fmt.Sprintf("Student: %s, is now registered for %s", name, className)
}

func (s *Student) RegisterSubject(name string, className string) string {
	return fmt.Sprintf("Student: %s, is now registered for %s class", name, className)
}

func (s *Student) DropClass(courses *[]map[Student]string, classname string) *[]map[Student]string {
	for i, course := range *courses {
		if _, exists := course[*s]; exists && course[*s] == classname {// first checks if the course exists, then checks if it exists and if it is equal to the classname to be dropped
			*courses = append((*courses)[:i], (*courses)[i+1:]...) // course[*s] == classname ensures that the subject associated with the student matches the provided classname
		 // and then the append function, takes all the elements up until the present element and then everything after it, and appends skipping the present element.
		}
		// course[*s] is used 
	}
	return courses
}

func (s *Student) DropSubject(subjects *[]map[Student]string, subjectname string) *[]map[Student]string {
	// loop through subjects
	for i, subject := range *subjects {
		// check if subject exists, and then if any subjects match the given subject name
		if _, exists := subject[*s]; exists && subject[*s] == subjectname{ // explore using the inbuilt map delete function to remove instead of  the append
			*subjects = append((*subjects)[:i],(*subjects)[i+1:]... ) // we want to remove the found subject at element i, that matches the condition

		}
	}
	return subjects
	//subject[*s] works because to access the value of the map, we call the key value
	
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
	courseToBeDropped := newStudent.DropClass(&courses, "A")
	fmt.Println("This is the course to be dropped:", courseToBeDropped)
	
	fmt.Println(newStudent)
	fmt.Println(courses)
}
// I think it's going to rain today 