package model

import "time"

type Workshop struct {
	Course // Embedded struct
	Instructor
	Date time.Time
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}

func (c Workshop) SignUp() bool {
	return true
}
