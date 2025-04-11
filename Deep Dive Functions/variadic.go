package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	// first parameter (20) is a starting int value, others are slice {10, 15, 40, -5}
	sum := sumup(20, 10, 15, 40, -5)
	anotherSum := sumup(20, numbers...)
	
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}