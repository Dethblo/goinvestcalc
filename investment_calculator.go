package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// basic approach
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)

	// formatted string approach
	//fmt.Printf("Future Value: %.2f\n", futureValue)
	//fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	// multi-line approach
	fmt.Printf(`
Future Value: %.2f
Future Value (adjusted for inflation): %.2f
	`, futureValue, futureRealValue)
}
