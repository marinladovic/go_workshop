package main

// global variables
var url = "https://frontendmasters.com"

func main() {
	// function-scoped variables
	const maxSpeed byte = 60

	message := "Hello from Go!"
	price := 34.4

	print(message, price, maxSpeed, url)	
}