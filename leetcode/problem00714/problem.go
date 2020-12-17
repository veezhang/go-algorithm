package problem00714

func maxProfit1(prices []int, fee int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}
	maxProfitValue := 0
	// 在最大化收益的前提下，如果我们手上拥有一支股票，那么它的最低买入价格是多少
	buy := prices[0] + fee // 买入价格 = 股票价格 + 手续费
	for i := 1; i < length; i++ {
		if prices[i]+fee < buy { // 如果可以以更低的价格买入，则更新
			buy = prices[i] + fee
		} else if prices[i] > buy { // 如果现在股票的价格 > buy ，则卖掉
			// 如果当前的股票价格 prices[i] 大于 buy，那么我们直接卖出股票并且获得 prices[i]−buy 的收益。
			// 但实际上，我们此时卖出股票可能并不是全局最优的（例如下一天股票价格继续上升），因此我们可以
			// 提供一个反悔操作，看成当前手上拥有一支买入价格为 prices[i] 的股票，将 buy 更新为 prices[i]。
			// 这样一来，如果下一天股票价格继续上升，我们会获得 prices[i+1]−prices[i] 的收益，加上这一天
			// prices[i]−buy 的收益，恰好就等于在这一天不进行任何操作，而在下一天卖出股票的收益
			maxProfitValue += prices[i] - buy
			buy = prices[i]
		} // 其它情况[buy−fee,buy] ，没必要买，也没必要卖
	}

	return maxProfitValue
}

func maxProfit2(prices []int, fee int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}

	dp := make([][2]int, length)
	// dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]
}

func maxProfit3(prices []int, fee int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < length; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]-fee), max(dp1, dp0-prices[i])
	}

	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
