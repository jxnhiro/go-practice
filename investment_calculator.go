package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years Kept: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFutureValues := fmt.Sprintf(`Future Value: %.1f
	Future Value (adjusted for inflation): %.1f`,futureValue, futureRealValue)

	fmt.Print(formattedFutureValues)
}

// A function name should define what a function does
func outputText(text string) {
	fmt.Print(text)
}