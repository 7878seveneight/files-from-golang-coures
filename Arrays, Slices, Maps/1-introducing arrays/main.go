package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}



func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	// from second elemtnt(included) to the third element (forth is excluded)
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)
	featuredPrices[0] = 100.99

	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// cant start with -1 (last index of an array)
	// featuredPrices = prices[-1:]

	fmt.Print("\n\n\n\n")
	
	newPrices := []float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(newPrices)

	discountPrices := []float64{101, 88.99, 20.59}
	// *listWithNewElements*... adds separate elements to the list
	// cant add *listWithNewElements* by itself because the main list isn't a list of lists
	newPrices = append(newPrices, discountPrices...)

	fmt.Println(newPrices)
}

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99

// 	prices = append(prices, 5.99)
// 	fmt.Println(prices)
// }