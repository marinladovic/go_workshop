package data

func CalculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func CalculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

func Birthday(age *int) {
	if *age > 35 {
		panic("Too old to be true")
	}
	*age++
}
