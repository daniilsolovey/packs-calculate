package pack

import (
	"fmt"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	calculator := &PCalculator{}

	tests := []struct {
		order          int
		expectedResult map[int]int
	}{
		{1, map[int]int{250: 1}},
		{250, map[int]int{250: 1}},
		{251, map[int]int{500: 1}},
		{500, map[int]int{500: 1}},
		{502, map[int]int{500: 1, 250: 1}},
		{800, map[int]int{1000: 1}},
		{1000, map[int]int{1000: 1}},
		{1200, map[int]int{1000: 1, 250: 1}},
		{1500, map[int]int{1000: 1, 500: 1}},
		{2000, map[int]int{2000: 1}},
		{2100, map[int]int{2000: 1, 250: 1}},
		{2100, map[int]int{2000: 1, 250: 1}},
		{2500, map[int]int{2000: 1, 500: 1}},
		{12001, map[int]int{250: 1, 2000: 1, 5000: 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Order %d", test.order), func(t *testing.T) {
			result, err := calculator.CalculatePacks(test.order, []int{5000, 2000, 1000, 500, 250})
			if err != nil {
				t.Fatal(err)
			}

			if !compareResults(result, test.expectedResult) {
				t.Errorf("Expected %v, got %v", test.expectedResult, result)
			}
		})
	}
}

func compareResults(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
