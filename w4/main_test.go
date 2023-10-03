package main

import (
	"errors"
	"testing"
)

func TestCanBuyCar(t *testing.T) {
	testCases := []struct {
		Name           string
		Age            uint8
		Salary         uint64
		ExpectedResult bool
		ExpectedError  error
	}{
		{Name: "John", Age: 120, Salary: 0, ExpectedResult: false, ExpectedError: InvalidAgeError},
		{Name: "Alice", Age: 17, Salary: 1000000, ExpectedResult: false, ExpectedError: nil},
		{Name: "Tom", Age: 30, Salary: 26000000, ExpectedResult: true, ExpectedError: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := CanBuyCar(Customer{Name: tc.Name, Age: tc.Age, Salary: tc.Salary})
			if result != tc.ExpectedResult {
				t.Errorf("Expected result: %v, but got: %v", tc.ExpectedResult, result)
			}
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("Expected error: %v, but got: %v", tc.ExpectedError, err)
			}
		})
	}
}
