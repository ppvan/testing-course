package main

import "errors"

type Customer struct {
	Name   string
	Age    uint8
	Salary uint64
}

var InvalidAgeError = errors.New("Invalid age")

func CanBuyCar(customer Customer) (bool, error) {

	if customer.Age > 100 {
		return false, InvalidAgeError
	}

	if customer.Age < 18 || customer.Salary < 25000000 {
		return false, nil
	}

	return true, nil
}

func main() {
	// ...
}
