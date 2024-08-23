package main

import "fmt"

func main() {
	var ptrAge *int
	age := 25

	ptrAge = &age

	fmt.Println(ptrAge)
	fmt.Println(*ptrAge)

	getAdultYears(ptrAge)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
