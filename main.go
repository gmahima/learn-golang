package main

import (
	"fmt"
	"learn-golang/quicksort"
)

func main() {
	fmt.Println("Hello, World! Welcome to my Go learning journey!")
	numbers := []int{10, 3, 1, 8, 10, 6, 2, 1}
	sortedNumbers := quicksort.Sort(numbers)
	fmt.Println("Sorted numbers:", sortedNumbers)
} 