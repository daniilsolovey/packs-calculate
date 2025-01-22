package main

import (
	"fmt"
	"log"

	"github.com/daniilsolovey/packs-calculate/internal/pack"
)

func main() {
	number := 800
	packs := []int{5000, 2000, 1000, 500, 250}

	result, err := pack.CalculatePacks(number, packs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", result)
}
