package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "USA"

	Prices := [2]int{100, 150}

	qty := len(Countries)

	// Slices
	names := []string{"Mary", "John"}
	names = append(names, "Carol")

	// Maps
	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	fmt.Println("Countries saved")
	fmt.Println(qty)
	fmt.Println(Prices)
	fmt.Println(names)
	fmt.Println(wellKnownPorts)
}
