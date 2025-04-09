package lists

import "fmt"

func main() {
	prices := []float64{10.99, 9.99}
	fmt.Println(prices)

	discountPrices := []float64{101, 88.99, 20.59}
	// *listWithNewElements*... adds separate elements to the list
	// cant add *listWithNewElements* by itself because the main list isn't a list of lists
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}