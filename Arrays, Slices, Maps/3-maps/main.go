package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// sec 7: lec 122. make()
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.7

	courseRating.output()

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRating {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}