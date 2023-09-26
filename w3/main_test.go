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
		// Test case 1: 0 <= Age < 18, Salary < 25m, ExpectedResult: false, ExpectedError: nil
		{Name: "John", Age: 10, Salary: 20000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 2: 18 <= Age <= 100, Salary < 25m, ExpectedResult: false, ExpectedError: nil
		{Name: "Alice", Age: 25, Salary: 15000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 3: Age > 100, Salary < 25m, ExpectedResult: false, ExpectedError: nil
		{Name: "Elderly", Age: 110, Salary: 20000000, ExpectedResult: false, ExpectedError: InvalidAgeError},

		// Test case 4: 0 <= Age < 18, 25m <= Salary <= 50m, ExpectedResult: false, ExpectedError: nil
		{Name: "Young Rich", Age: 15, Salary: 35000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 5: 18 <= Age <= 100, 25m <= Salary <= 50m, ExpectedResult: false, ExpectedError: nil
		{Name: "Adult Rich", Age: 45, Salary: 45000000, ExpectedResult: true, ExpectedError: nil},

		// Test case 6: Age > 100, 25m <= Salary <= 50m, ExpectedResult: false, ExpectedError: nil
		{Name: "Elderly Rich", Age: 105, Salary: 48000000, ExpectedResult: false, ExpectedError: InvalidAgeError},

		// Test case 7: 0 <= Age < 18, Salary > 50m, ExpectedResult: true, ExpectedError: nil
		{Name: "Young Millionaire", Age: 12, Salary: 75000000, ExpectedResult: false, ExpectedError: nil},
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

func TestCanBuyCarEquivalent(t *testing.T) {
	testCases := []struct {
		Name           string
		Age            uint8
		Salary         uint64
		ExpectedResult bool
		ExpectedError  error
	}{
		// Test case 1: Age in [0, 18), Salary in [-infinity, 25m), ExpectedResult: false, ExpectedError: nil
		{Name: "Teenager Low Income", Age: 10, Salary: 15000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 2: Age in [0, 18), Salary in [25m, 50m], ExpectedResult: false, ExpectedError: nil
		{Name: "Teenager Middle Income", Age: 17, Salary: 35000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 3: Age in [0, 18), Salary in [50m, infinity), ExpectedResult: true, ExpectedError: nil
		{Name: "Teenager High Income", Age: 16, Salary: 75000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 4: Age in [18, 100], Salary in [-infinity, 25m), ExpectedResult: false, ExpectedError: nil
		{Name: "Adult Low Income", Age: 25, Salary: 15000000, ExpectedResult: false, ExpectedError: nil},

		// Test case 5: Age in [18, 100], Salary in [25m, 50m], ExpectedResult: true, ExpectedError: nil
		{Name: "Adult Middle Income", Age: 45, Salary: 35000000, ExpectedResult: true, ExpectedError: nil},

		// Test case 6: Age in [18, 100], Salary in [50m, infinity), ExpectedResult: false, ExpectedError: nil
		{Name: "Adult High Income", Age: 75, Salary: 75000000, ExpectedResult: true, ExpectedError: nil},

		// Test case 7: Age in [100, infinity), Salary in [-infinity, 25m), ExpectedResult: false, ExpectedError: nil
		{Name: "Elderly Low Income", Age: 105, Salary: 15000000, ExpectedResult: false, ExpectedError: InvalidAgeError},

		// Test case 8: Age in [100, infinity), Salary in [25m, 50m], ExpectedResult: false, ExpectedError: InvalidAgeError
		{Name: "Elderly Middle Income", Age: 110, Salary: 35000000, ExpectedResult: false, ExpectedError: InvalidAgeError},

		// Test case 9: Age in [100, infinity), Salary in [50m, infinity), ExpectedResult: false, ExpectedError: InvalidAgeError
		{Name: "Elderly High Income", Age: 120, Salary: 75000000, ExpectedResult: false, ExpectedError: InvalidAgeError},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := CanBuyCar(Customer{Name: tc.Name, Age: tc.Age, Salary: tc.Salary})
			if result != tc.ExpectedResult {
				t.Errorf("Expected result: %v, but got: %v", tc.ExpectedResult, result)
			}
			if err != tc.ExpectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.ExpectedError, err)
			}
		})
	}
}
