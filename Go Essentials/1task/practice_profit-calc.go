package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("revenue: ")

	// fmt.Print("expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInput("expenses: ")

	// fmt.Print("tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate := getUserInput("tax rate(%): ")

	earningsBeforeTax, profit, ratio := calcValues(revenue, expenses, taxRate)
	
	fmt.Printf("\nEBT: %.1f\n", earningsBeforeTax)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.3f\n", ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calcValues(revenue, expenses, taxRate float64) (earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ratio = earningsBeforeTax / profit
	return
} 
