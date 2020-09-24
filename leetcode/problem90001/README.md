# random5 generate random7

## 题目

给你一个能生成1到5随机数的函数，用它写一个函数生成1到7的随机数

## 解题思路

### 二进制法

1,2,3,4,5
1,2 转为 0
3 重做
4,5 转为 1

然后 0,1 来随机 1-7 (7个数)，则需要 3 个二进制来计算，则随机生成 0 - 7（8个数），7 的话重做

### N进制法

上面需要转换，其实我们这里可以使用 5 进制来处理，要生产 1-7 的数字，可以使用：

`random5` 生成`[1,5]`的数，可以转化为`[0,5)`的数，`5 * random5 + random5` 就可以等概率生成`[0,25)`的数，那么 `[0,21)`均分给`7`个数，`[21,25)重做`

这里实现了一个通用的：

```go
// go run -v ./leetcode/problem90001/main
type Random struct {
	Min int
	Max int
}

// 生成随机数[rb.Min, rb.Max]，根据 random 来计算Random
func (rb *Random) random() int {
	// rb.Min + [0, rb.Max-rb.Min+1) => [rb.Min, rb.Max+1)
	return rb.Min + rand.Intn(rb.Max-rb.Min+1)
}

// 生成随机数[min, max]，根据 random 来计算Random
func (rb *Random) Random(min, max int) int {
	// 比如 rb.Min=1,rb.Max=10 ，则可以生成10个数， base = rb.Max - rb.Min + 1 = 10
	// 假设要生产 [10,120]之间的随机数，则需要生成 120 - 10 + 1 = 111 个随机数
	// 至少需要 3 次随机，才能使得随机数大于 111
	// 100 * [0, 10) + 10 * [0, 10) + [0, 10) 则生成 [0, 1000) 之间的随机数，总共 1000 个
	// 如果把 [111, 1000)都丢弃，会需要更多的尝试，所以，可以最大倍数进行映射
	// 比如对于 [0,999) 直接的数对于 111 映射，如果结果是 [999,1000) 则重新随机
	// for ( randomNumber >= 1000 - 1000 % 111 ) {
	// 		randomNumber = 100 * rand10() + 10 * rand10() + rand10()
	// }
	// 最终返回的是： randomNumber % 111 + 10 ，对应的也就是： [10, 120]

	// 进制，按照这个进制来计算
	base := rb.Max - rb.Min + 1
	// 需要的随机数字个数
	randomNeeded := max - min + 1

	// 计算需要多少位base进制，才能使结果大于 n
	// 假设需要 baseBit 位的 base 进制，则可以随机 pow(base, baseBit) 个数字
	baseBit := 1
	randomAll := base
	// 用来记录每一位需要乘以的基数
	baseTable := []int{1}

	for randomAll < randomNeeded {
		baseBit++
		baseTable = append(baseTable, randomAll)
		randomAll *= base
	}

	// 可用的随机数
	randomAvailable := randomAll - randomAll%randomNeeded
	randomNumber := randomAvailable // 初始化一个 >= randomAvailable 的数字
	randomCount := 0                // 记录随机次数
	for randomNumber >= randomAvailable {
		if randomNumber >= randomAvailable {
			randomCount++
		}
		randomNumber = 0
		// 生成随机数
		for i := 0; i < baseBit; i++ {
			// rb.random() - rb.Min ==> [rb.Min, rb.Max] ==>[0, rb.Max - rb.Min] ==> [0, base)
			randomNumber += (rb.random() - rb.Min) * baseTable[i]
		}
	}

	if false {
		fmt.Printf("[%d,%d]:%d => [%d,%d]:%d baseBit: %d, randomAll: %d, randomAvailable: %d randomNumber: %d randomCount: %d, return: %d\n",
			rb.Min, rb.Max, base,
			min, max, randomNeeded,
			baseBit, randomAll,
			randomAvailable, randomNumber, randomCount,
			randomNumber%randomNeeded+min,
		)
	}

	return randomNumber%randomNeeded + min
}
```
