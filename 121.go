package main

func maxProfit(prices []int) int {
    maxpro := 0
	mini := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < mini {
			mini = prices[i]
		}
		profit := prices[i] - mini
		if profit > maxpro {
			maxpro = profit
		}
	}
	return maxpro
}