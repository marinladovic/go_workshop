package main

import (
	"fmt"

	"marinladovic.com/go/workshop/data"
)

var name = "Frontend Masters"

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func main() {
	fmt.Print("dfdfd")
	fmt.Print(data.MaxSpeed)

	PrintData()
}
