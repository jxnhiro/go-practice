package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years Kept: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future Value (adjusted for inflation): %.1f", futureRealValue)

	fmt.Print(formattedFutureValue, formattedFutureRealValue)
}