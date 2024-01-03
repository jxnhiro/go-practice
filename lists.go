package main

type Product struct {
	title string
	id string
	price float64
}

func main() {
	prices := []float64{}
	
}
// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices[2])

// 	//Slices are new arrays
// 	pricesSecondToEnd := prices[2:]
// 	pricesStopInThird := prices[:3]
// 	fmt.Println(pricesSecondToEnd)
// 	fmt.Println(pricesStopInThird)

// 	//We can create slices from slices
// 	pricesSubset := pricesStopInThird[:2]
// 	fmt.Println(pricesSubset)

// 	//You can reslice to the right but not to the start
// 	featuredPrices := prices[1:2]
// 	fmt.Println(featuredPrices)
// 	featuredPrices = featuredPrices[:3]
// 	fmt.Println(featuredPrices)
// }