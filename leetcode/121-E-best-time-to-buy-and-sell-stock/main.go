package main

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	return a + b - getMin(a, b)
}

func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices < 2 {
		return 0
	}

	histNumMin := getMin(prices[0], prices[1])
	histProfMax := prices[1] - prices[0]
	for k := 1; k < lenPrices; k++ {
		histProfMax = getMax(histProfMax, prices[k]-histNumMin)
		histNumMin = getMin(histNumMin, prices[k])
	}
	// fmt.Println(histProfMax)
	return getMax(0, histProfMax)
}

func main() {

}
