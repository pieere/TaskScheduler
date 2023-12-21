package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 2, 8, 1, 6}

	sort.Slice(numbers, func(i, j int) bool {
		fmt.Print(i, j)
		fmt.Print("________________________")
		return numbers[i] < numbers[j]
	})

	fmt.Println("Sorted Numbers:", numbers)
}

