package problem00122

func maxProfit1(prices []int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}
	maxProfitValue := 0

	for i := 1; i < length; {
		// 先找一个最小值，下降阶段
		for ; i < length && prices[i-1] >= prices[i]; i++ {
		}
		buy := prices[i-1]
		// 再找一个最大值，上升阶段
		for ; i < length && prices[i-1] <= prices[i]; i++ {
		}
		maxProfitValue += prices[i-1] - buy
	}
	return maxProfitValue
}

func maxProfit2(prices []int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}
	maxProfitValue := 0
	for i := 1; i < length; i++ {
		maxProfitValue += max(0, prices[i]-prices[i-1])
	}

	return maxProfitValue
}

func maxProfit3(prices []int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}

	dp := make([][2]int, length)
	// dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]
}

func maxProfit4(prices []int) int {
	length := len(prices)
	if 1 >= length {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < length; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}

	return dp0
}

// 单调栈
func maxProfit5(prices []int) int {
	if 1 >= len(prices) {
		return 0
	}

	stack := []int{}
	maxProfitValue := 0

	for _, price := range prices {
		// 如果栈不为空，且栈顶元素大于等于 price
		for len(stack) > 0 && stack[len(stack)-1] >= price {
			maxProfitValue += stack[len(stack)-1] - stack[0]
			stack = stack[:0] // 清空栈
		}
		// 入栈
		stack = append(stack, price)
	}

	// 如果栈里面还有元素，可能需要更新下 maxProfitValue
	if len(stack) > 0 {
		maxProfitValue += stack[len(stack)-1] - stack[0]
	}

	return maxProfitValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
