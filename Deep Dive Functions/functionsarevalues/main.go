package functionsarevalues

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	triple := transformNumbers(&numbers, triple)
	fmt.Println(triple)

	moreNumbers := []int{5, 1, 2}
	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) func(int) int{
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
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