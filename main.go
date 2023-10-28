package main

import (
	"fmt"

	"marinladovic.com/go/workshop/data"
)

func main() {
	_, cityTax := data.CalculateTaxWithName(75)
	fmt.Println(cityTax)

	defer fmt.Println("Bye!!")

	age := 31
	data.Birthday(&age)
	data.Birthday(&age)
	fmt.Printf("age: %v\n", age)
}
