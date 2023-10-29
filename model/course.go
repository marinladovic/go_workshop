package model

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

// Define a method that will be the string representation of the struct
// Stringer interface
func (c Course) String() string {
	return fmt.Sprintf("--- %v --- (%v)", c.Name, c.Instructor.FirstName)
}

func (c Course) SignUp() bool {
	return true
}
