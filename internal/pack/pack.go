package pack

import (
	"errors"
	"math"
	"sort"
)

type PackCalculator interface {
	CalculatePacks(number int, packs []int) (map[int]int, error)
}

// PCalculator implements the PackCalculator interface.
type PCalculator struct{}

func NewPackCalculator() *PCalculator {
	return &PCalculator{}
}

// CalculatePacks calculates the optimal number of packs for a given order.
func (c *PCalculator) CalculatePacks(number int, packs []int) (map[int]int, error) {
	if number <= 0 {
		return nil, errors.New("number must be greater than zero")
	}

	if len(packs) == 0 {
		return nil, errors.New("packs list cannot be empty")
	}

	// calculate the maximum numbers possible in result
	maxValues := number + getMinimumPack(packs) + 1
	temp := maxValues

	// initialize the dynamicArray array and set values with 0
	dynamicArray := make([]int, maxValues)
	dynamicArray[0] = 0
	for i := 1; i < maxValues; i++ {
		dynamicArray[i] = -1
	}

	// add values to dynamic array based on the packs
	for _, i := range packs {
		for j := i; j < maxValues; j++ {
			if dynamicArray[j-i] == -1 {
				continue
			}

			if dynamicArray[j] == -1 {
				dynamicArray[j] = dynamicArray[j-i] + 1
			} else {
				dynamicArray[j] = int(
					math.Min(
						float64(dynamicArray[j]),
						float64(dynamicArray[j-i]+1),
					),
				)
			}

			if j >= number {
				temp = int(
					math.Min(float64(temp),
						float64(j)),
				)
			}
		}
	}

	// create the result map
	result := make(map[int]int)
	for temp != 0 {
		for _, i := range packs {
			if temp >= i && dynamicArray[temp-i]+1 == dynamicArray[temp] {
				temp = temp - i
				result[i]++
				break
			}
		}
	}

	return result, nil
}

func getMinimumPack(data []int) int {
	sort.Ints(data)
	return data[0]
}
