package main

import (
	"fmt"

	"marinladovic.com/go/workshop/model"
)

func main() {
	max := model.Instructor{Id: 3, LastName: "Firtman"}
	max.FirstName = "Maximiliano"

	kyle := model.NewInstructor("Kyle", "Simmons")

	goCourse := model.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}

	fmt.Printf("%+v\n", goCourse)

	kyle.Print()
}
