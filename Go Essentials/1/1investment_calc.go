package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 2.5

	// fmt.Print("Your investment amount: ")
	outputText("Your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Your expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Your years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValues(investmentAmount, expectedReturnRate, years)
	
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value(with inflation: %.1f\n)", futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Printf("Future Value: %.1f Future Value(with inflation: %.1f)", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calcFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}