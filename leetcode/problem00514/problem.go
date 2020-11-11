package problem00514

import (
	"math"
)

func findRotateSteps1(ring string, key string) int {
	// 统计每个字符的位置列表
	posMap := [26][]int{}
	for i, ch := range ring {
		// 肯定是递增的
		posMap[ch-'a'] = append(posMap[ch-'a'], i)
	}

	// 初始化
	lenRing, lenKey := len(ring), len(key)
	// 定义 `dp[i][j]` 表示初始 `ring` 指向 `j` 的情况下，要匹配 `key[i]` 和之前所有字母所需的最少旋转次数
	// 也就是 `key[0-i]` 都匹配好，需要最少旋转次数
	dp := make([][]int, lenKey)
	for i := range dp {
		dp[i] = make([]int, lenRing)
	}
	// 最少旋转次数
	m := math.MaxInt64

	// 初始状态，计算 key[0] 对齐，需要旋转次数
	for _, j := range posMap[key[0]-'a'] {
		// 当前处于 j 位置， 移到 0 位置需要多少步骤
		dp[0][j] = minRotate(j, 0, lenRing)
		if 1 == lenKey { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
			m = min(m, dp[0][j])
		}
	}

	// 然后后向推算其他的所有对齐需要多少步
	for i := 1; i < lenKey; i++ {
		// 计算对字符 key[i] 的每个位置(posMap[key[i]-'a']) ，
		// 从前一个字符 key[i-1] 的每个位置(posMap[key[i-1]-'a']) 旋转而来，最少需要旋转次数
		if i == lenKey-1 { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
			m = math.MaxInt64
		}
		for _, j := range posMap[key[i]-'a'] {
			mij := math.MaxInt64
			for _, k := range posMap[key[i-1]-'a'] {
				mij = min(mij, dp[i-1][k]+minRotate(j, k, lenRing))
			}
			dp[i][j] = mij
			if i == lenKey-1 { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
				m = min(m, mij)
			}
		}
	}
	// 获取到最后一个字符需要最少旋转次数 + 按几下
	return m + lenKey
}

func findRotateSteps2(ring string, key string) int {
	// 统计每个字符的位置列表
	posMap := [26][]int{}
	for i, ch := range ring {
		// 肯定是递增的
		posMap[ch-'a'] = append(posMap[ch-'a'], i)
	}

	// 初始化
	lenRing, lenKey := len(ring), len(key)
	// 定义 `dp[i][j]` 表示初始 `ring` 指向 `j` 的情况下，要匹配 `key[i]` 和之前所有字母所需的最少旋转次数
	// 也就是 `key[0-i]` 都匹配好，需要最少旋转次数
	// 空间优化，dp[0] 表示前一个， dp[0] 表示后一个，
	dp := [2][]int{}
	for i := range dp {
		dp[i] = make([]int, lenRing)
	}
	// 最少旋转次数
	m := math.MaxInt64

	// 初始状态，计算 key[0] 对齐，需要旋转次数
	for _, j := range posMap[key[0]-'a'] {
		// 当前处于 j 位置， 移到 0 位置需要多少步骤
		dp[0][j] = minRotate(j, 0, lenRing)
		if 1 == lenKey { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
			m = min(m, dp[0][j])
		}
	}

	// 然后后向推算其他的所有对齐需要多少步
	for i := 1; i < lenKey; i++ {
		// 计算对字符 key[i] 的每个位置(posMap[key[i]-'a']) ，
		// 从前一个字符 key[i-1] 的每个位置(posMap[key[i-1]-'a']) 旋转而来，最少需要旋转次数
		if i == lenKey-1 { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
			m = math.MaxInt64
		}
		for _, j := range posMap[key[i]-'a'] {
			mij := math.MaxInt64
			for _, k := range posMap[key[i-1]-'a'] {
				mij = min(mij, dp[0][k]+minRotate(j, k, lenRing))
			}
			dp[1][j] = mij
			if i == lenKey-1 { // 如果是 key 中最后一个，统计下最小旋转次数，不用初始化 dp 为 math.MaxInt64
				m = min(m, mij)
			}
		}
		dp[1], dp[0] = dp[0], dp[1]
	}
	// 获取到最后一个字符需要最少旋转次数 + 按几下
	return m + lenKey
}

// minRotate 返回长度为 length 的循环数组中，从 from 到 to 需要的步数（可以前后两边走）
func minRotate(from, to, length int) int {
	r := (from - to + length) % length
	if r > length-r {
		return length - r
	}
	return r
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
