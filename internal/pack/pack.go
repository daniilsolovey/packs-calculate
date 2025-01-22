package pack

import (
	"errors"
	"sort"
)

// CalculatePacks calculates the optimal number of packs for a given order.
func CalculatePacks(number int, packs []int) (map[int]int, error) {
	if number <= 0 {
		return nil, errors.New("number must be greater than zero")
	}

	if len(packs) == 0 {
		return nil, errors.New("packs list cannot be empty")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	result := make(map[int]int)

	for _, packSize := range packs {
		if number >= packSize {
			count := number / packSize
			result[packSize] = count
			number -= count * packSize
		}
	}

	if number > 0 {
		smallestPack := packs[len(packs)-1]
		result[smallestPack]++
		number = 0
	}

	return optimizeResult(result, packs)
}

// optimizeResult tries to consolidate smaller packs into a larger one.
func optimizeResult(result map[int]int, packs []int) (map[int]int, error) {
	optimized := false
	for packSize, count := range result {
		for _, largerPack := range packs {
			if packSize < largerPack && count >= 2 {
				if packSize*count >= largerPack {
					result[largerPack]++
					result[packSize] -= 2
					if result[packSize] <= 0 {
						delete(result, packSize)
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
