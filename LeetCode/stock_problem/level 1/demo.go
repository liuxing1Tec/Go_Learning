package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	gaps := make([]int, len(prices)-1)
	for i := 0; i < len(gaps); i++ {
		gaps[i] = prices[i+1] - prices[i]
	}
	ans := 0
	sum := 0
	for i := 0; i < len(gaps); i++ {
		sum += gaps[i]
		ans = getMax(ans, sum)
		sum = getMax(sum, 0)
	}
	return ans
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
