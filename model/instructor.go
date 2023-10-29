package model

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// Factory function
func NewInstructor(firstName, lastName string) Instructor {
	return Instructor{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (i Instructor) Print() {
	fmt.Printf("%s %s (%d)\n", i.FirstName, i.LastName, i.Score)
}
