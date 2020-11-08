package problem00121

import (
	"math"
)

func maxProfit1(prices []int) int {
	if 1 >= len(prices) {
		return 0
	}

	// 到目前为止的最小价格
	minPriceValue := math.MaxInt64
	// 到目前为止的最大数收益
	maxProfitValue := 0

	for _, price := range prices {
		// 更新到目前为止的最小价格
		// 更新的时候，不需要计算 maxProfitValue 了， 因为当天看成买入的话，并不能获取更多的利益
		if price < minPriceValue {
			minPriceValue = price
		} else if price-minPriceValue > maxProfitValue {
			maxProfitValue = price - minPriceValue
		}
	}
	return maxProfitValue
}

func maxProfit2(prices []int) int {
	if 1 >= len(prices) {
		return 0
	}

	leftPointer := 0
	rightPointer := 1
	maxProfitValue := 0

	for rightPointer < len(prices) {
		if prices[rightPointer] < prices[leftPointer] {
			leftPointer = rightPointer
			rightPointer++
			continue
		}
		maxProfitValue = max(maxProfitValue, prices[rightPointer]-prices[leftPointer])
		rightPointer++
	}
	return maxProfitValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit3(prices []int) int {
	if 1 >= len(prices) {
		return 0
	}

	// 到目前为止的最小价格
	minPriceValue := math.MaxInt64
	dp := make([]int, len(prices))

	for i, price := range prices {
		// 更新到目前为止的最小价格
		if price < minPriceValue {
			minPriceValue = price
		}
		// 更新 dp 数组
		// 其实这里 dp 数组优化，就是方法1
		if i >= 1 {
			dp[i] = max(dp[i-1], price-minPriceValue)
		}
	}
	return dp[len(prices)-1]
}

// 单调栈
func maxProfit4(prices []int) int {
	if 1 >= len(prices) {
		return 0
	}

	stack := []int{}
	maxProfitValue := 0

	// prices = append(prices, -1)

	for _, price := range prices {
		// 如果栈不为空，且栈顶元素大于等于 price
		for len(stack) > 0 && stack[len(stack)-1] >= price {
			if stack[len(stack)-1]-stack[0] > maxProfitValue {
				maxProfitValue = stack[len(stack)-1] - stack[0]
			}

			stack = stack[:len(stack)-1]
		}
		// 入栈
		stack = append(stack, price)
	}

	// 如果栈里面还有元素，可能需要更新下 maxProfitValue
	if len(stack) > 0 {
		if stack[len(stack)-1]-stack[0] > maxProfitValue {
			maxProfitValue = stack[len(stack)-1] - stack[0]
		}
	}

	return maxProfitValue
}
