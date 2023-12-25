package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years Kept: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	formattedFutureValues := fmt.Sprintf(`Future Value: %.1f
	Future Value (adjusted for inflation): %.1f`,futureValue, futureRealValue)

	fmt.Print(formattedFutureValues)
}

// A function name should define what a function does
func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	realFutureValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	return futureValue, realFutureValue
}