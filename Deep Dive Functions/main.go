package main

import "fmt"

type intFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	triple := transformNumbers(&numbers, triple)
	fmt.Println(triple)
}

func transformNumbers(numbers *[]int, transform intFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(num int) int {
	return num * 2
}

// func tripleNumbers(numbers *[]int) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, triple(val))
// 	}

// 	return dNumbers
// }
func triple(num int) int {
	return num * 3
}