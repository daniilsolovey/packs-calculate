package main

import (
	"errors"
	"fmt"
	"sort"
)

// calculatePacks function to calculate the optimal number of packs for a given order
func calculatePacks(number int, packs []int) (map[int]int, error) {
	if number <= 0 {
		return nil, errors.New("number must be greater than zero")
	}

	if len(packs) == 0 {
		return nil, errors.New("packs list cannot be empty")
	}

	// sort pack sizes in descending order for greedy approach
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))

	result := make(map[int]int)
	for _, packSize := range packs {
		if number >= packSize {
			count := number / packSize // How many of this pack we need
			result[packSize] = count   // Store the count of the current pack
			number -= count * packSize // Subtract the fulfilled items from the total number
		}
	}

	// if there are any items left, try to fulfill them with the smallest pack size
	if number > 0 {
		// ensure there is a pack that can fulfill the remaining number of items
		smallestPack := packs[len(packs)-1]
		result[smallestPack]++
		number = 0 // fulfilled the remaining items
	}

	return optimizeResult(result, packs)

}

// Recursive function to optimize the result map by replacing smaller packs with larger ones
func optimizeResult(result map[int]int, packs []int) (map[int]int, error) {
	optimized := false

	// need to go through the result map and try to consolidate smaller packs into larger ones
	for packSize, count := range result {
		for _, largerPack := range packs {
			if packSize < largerPack && count >= 2 {
				// check if the combination of smaller packs can be replaced with a larger one
				if packSize*count >= largerPack {
					// consolidate the smaller packs into the larger pack
					result[largerPack]++
					result[packSize] -= 2 // replace 2 smaller packs with 1 larger pack
					if result[packSize] <= 0 {
						delete(result, packSize) // remove the smaller pack if count goes to zero
					}
					optimized = true
					break
				}
			}
		}
	}

	if optimized {
		return optimizeResult(result, packs)
	}

	return result, nil
}

func main() {
	var packSizes = []int{5000, 2000, 1000, 500, 250}

	number := 12001

	result, err := calculatePacks(number, packSizes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", result)
}
