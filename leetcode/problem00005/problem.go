package problem00005

// 最长公共子串
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func longestPalindromeLongestSubstring(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	// i 为原 字符串的下标
	// j 为反转后字符串的下标，也就对应原来的 length - 1 - j
	// dp[i][j] 为是回文长度，公共子串长度，加一些判断，便知道是否是回文
	// 如果s[i] == s[j]，则 dp[i][j] = 1 + dp[i-1][j-1]
	//
	//   ┌───────────────┬───────────────┐
	//   │ dp[i-1][j-1]  │               │
	//   ├───────────────┼───────────────┤
	//   │               │  dp[i][j]     │
	//   └───────────────┴───────────────┘
	//
	// dp[i][j] 只是依赖 dp[i-1][j-1] ， 只要确保 dp[i-1][j-1] 在 dp[i][j] 之前计算完成就好。
	// 可以画图看看是否满足以上依赖，这里只列举一种。
	//
	//	for i := 0; i < length; i++ {
	// 		for j := 0; j < length; j++ {
	//
	// 以上是竖着一排一排从前往后的来的，后面一排只需要前一排的数据，是否可以减少 dp ？
	// 设置dp为一维 dp := make([]int, length)，表示某一列的值，j 从 0 到 length-1 。
	// 如果 s[i] = s[j]
	// i 列    i-1 列
	// dp[0] = 1
	// dp[1] = dp[0] + 1
	// dp[2] = dp[1] + 1
	// dp[3] = dp[2] + 1
	// ...
	// 修改 dp[3] 的时候，dp[2] 已经被改动过，不再是前一竖排的数据，我们将 j 倒着来看看
	// 如果 s[i] = s[j]
	// i 列    i-1 列
	// ...
	// dp[3] = dp[2] + 1
	// dp[2] = dp[1] + 1
	// dp[1] = dp[0] + 1
	// dp[0] = 1
	//
	// 修改 dp[3] 的时候，dp[2] 还未被改动过，是前一竖排的数据，可以！
	//
	dp := make([]int, length)

	// maxLen 为最长字串的长度
	maxLen := 0
	endPos := 0

	for i := 0; i < length; i++ {
		for j := length - 1; j >= 0; j-- {
			// j 反转后字符串的下标， 也就对应原来的 length - 1 - j
			jBeforeReversePos := length - 1 - j
			if s[i] == s[jBeforeReversePos] {
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					dp[j] = 1 + dp[j-1]
				}

				if dp[j] > maxLen {
					// i 应该找到回文的最后一个字符， jBeforeReversePos 位置应该在前面
					// 并且 i 和 jBeforeReversePos 相差 dp[j] 也就是回文长度
					if jBeforeReversePos+dp[j]-1 == i {
						maxLen = dp[j]
						endPos = i
					}
				}
			} else {
				dp[j] = 0
			}
		}
	}

	return s[endPos-maxLen+1 : endPos+1]
}

// 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func longestPalindromeDP(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	// i 为回文起始位置，j 为回文结束位置
	// dp[i][j] = true 如果 s[i] ... s[j] 为回文，否则为 false
	// j >= i ， j == i 则是空串，dp[i][j] = true ；j == i + 1 则是单字符，dp[i][j] = true
	// 所以转换函数为： dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
	//
	//   ┌───────────────┬───────────────┐
	//   │               │  dp[i+1][j-1] │
	//   ├───────────────┼───────────────┤
	//   │   dp[i][j]    │               │
	//   └───────────────┴───────────────┘
	//
	// dp[i][j] 只是依赖 dp[i+1][j-1] ， 只要确保 dp[i+1][j-1] 在 dp[i][j] 之前计算完成就好。
	// 可以画图看看是否满足以上依赖，这里只列举一种。
	//
	//	for i := length - 1; i >= 0; i-- {
	// 		for j := 0; j < length; j++ {
	//
	// 以上是竖着一排一排从后往前的来的，前面一排只需要后一排的数据，是否可以减少 dp ？
	// 设置dp为一维 dp := make([]bool, length)，表示某一列的值，j 从 0 到 length-1 。
	// 如果 s[i] = s[j]
	// i 列    i+1 列
	// dp[0] = true （空串）
	// dp[1] = dp[0]
	// dp[2] = dp[1]
	// dp[3] = dp[2]
	// ...
	// 修改 dp[3] 的时候，dp[2] 已经被改动过，不再是前一竖排的数据，我们将 j 倒着来看看
	// 如果 s[i] = s[j]
	// i 列    i+1 列
	// ...
	// dp[3] = dp[2]
	// dp[2] = dp[1]
	// dp[1] = dp[0]
	// dp[0] = true （空串）
	//
	// 修改 dp[3] 的时候，dp[2] 还未被改动过，是后一竖排的数据，可以！
	//

	dp := make([]bool, length)
	// maxLen 为最长字串的长度
	// startPos记录起始下标
	maxLen := 0
	startPos := 0

	for i := length - 1; i >= 0; i-- {
		for j := length - 1; j >= i; j-- {
			if s[i] == s[j] {
				if j < i+3 {
					// 1. j = i 空串，为 true
					// 2. j = i + 1 , 单个字符， 为 true
					// 3. j = i + 2 , 双字符，并且 s[i] = s[j] ，即 AA类型，为 true
					// 3. j = i + 3 , 双字符，并且 s[i] = s[j] ，即 AxA类型，为 true
					dp[j] = true
				} else {
					dp[j] = dp[j-1]
				}

				// 由于我们这里是从后往前搜索的，因此，这里 = 的时候也需要修改，返回的结果应该是前面的
				if dp[j] && maxLen <= j-i+1 {
					maxLen = j - i + 1
					startPos = i
				}
			} else {
				dp[j] = false
			}
		}
	}

	return s[startPos : startPos+maxLen]
}

// 扩展中心
// 时间复杂度：O(n^2)，由于围绕中心来扩展回文会耗去 O(n)O(n) 的时间，所以总的复杂度为 O(n^2)
// 空间复杂度：O(1)
func longestPalindromeExandCenter(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	// maxLen 为最长字串的长度
	// startPos记录中间下标
	maxLen := 0
	centerPos := 0

	// 从 left， right 向两边扩展，直到不对成为止，返回长度
	exandCenter := func(s string, left, right int) int {
		for left >= 0 && right < length && s[left] == s[right] {
			left--
			right++
		}
		// left 到 right 长度是： right - left + 1
		// 但是此处是不满足的条件，left 已经减了1 ， right 已经加了1 ，所以最终需要 - 2
		// 即： right - left + 1 - 2
		return right - left - 1
	}

	for i := 0; i < length; i++ {
		// 回文是长度是奇数的时候
		evenLen := exandCenter(s, i, i)
		// 回文是长度是偶数的时候
		oddLen := exandCenter(s, i, i+1)

		// 获取一个大的，重用 evenLen
		if evenLen < oddLen {
			evenLen = oddLen
		}

		if maxLen < evenLen {
			maxLen = evenLen
			centerPos = i
		}
	}

	// 根据中心点计算起始位置，延迟到这里计算，只需要计算一次，重用下 centerPos
	// 比如 centerPos = 9, maxLen = 5, 则结果为： 7,8,9,10,11， 	9 - (5 - 1)/2
	// 比如 centerPos = 9, maxLen = 6, 则结果为： 7,8,9,10,11,12	9 - (6-  1)/2
	centerPos = centerPos - (maxLen-1)/2

	return s[centerPos : centerPos+maxLen]
}

// Manacher's Algorithm 马拉车算法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func longestPalindromeManacher(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	// 转换字符串，开头添加 ^ , 中间添加 # , 结尾添加 $ ，假定这3个字符不在字符串中
	// 比如： ab => ^#a#b#$ ;  aa => ^#a#a#$ ; aba => ^#a#b#a#$ ;
	// 转换后总长度一定为奇数， 回文长度也一定为奇数
	// 原回文长度=转换后的回文长度/2 ， maxLen = maxLen/2
	// 这里在开头和结尾加入的 ^, $ 符号，是为了中心扩展的时候，肯定不会超出边界，应为 ^, $ 不会等于任何字符
	sTrans := make([]byte, 2*length+3) // 2*length+1+2
	sTrans[0], sTrans[2*length+2] = '^', '$'
	for i := 0; i < length; i++ {
		sTrans[2*i+1] = '#'
		sTrans[2*i+2] = s[i]
	}
	sTrans[2*length+1] = '#'

	// length 为 sTrans长度了
	length = 2*length + 3

	// radius 为回文长度的半径（单边的长度）, center 为中心， right 为右边， left 为 right 对于 center 的对称点
	// 根据回文的对称性，根据左边的可以推断出右边的值
	// 如果 i 没有超过 right ，则 radius[i] = 对于 center 的对称点的 radius[center - (i - center)] ，即 radius[2*center - i)]
	// 否则， 此时 radius[i] = 0，因为右边的并不清楚
	// 如果过界了，需要扩展来计算
	//
	radius := make([]int, length)
	center := 0
	right := 0
	left := 0

	// maxLen 为转换后回文的半径（单边的长度）
	// startPos记录中间下标
	maxRadius := 0
	centerPos := 0

	for i := 1; i < length-1; i++ {
		if right > i {
			left = 2*center - i
			radius[i] = radius[left]
			// 可能超过 right 了， right 右边的情况我们并不清楚，所以最多只能到 right
			if radius[i] > right-i {
				radius[i] = right - i
			}
		}
		//else { radius[i] = 0 } 默认就是 0 ,不需要设置

		// 从 i 向两边扩展
		for sTrans[i+1+radius[i]] == sTrans[i-1-radius[i]] {
			radius[i]++

			// 只有这里才会增加 radius[i] ，上面最多也就是等于，所以在这里找最大值
			if maxRadius < radius[i] {
				maxRadius = radius[i]
				centerPos = i
			}
		}

		// 是否需要修改 right 边界
		if i+radius[i] > right {
			center = i
			right = i + radius[i]
		}
	}

	// maxRadius, centerPos 都为转换后的字符串的数值，这里需要转换为之前的，重用下 centerPos
	// maxRadius 为转换后的半径，正好等于原来的长度
	// 根据之前转换时候赋值 sTrans[2*i+2] = s[i] ， centerPos = 2*i+2 => i = (centerPos + 2)/2
	//
	// 比如： (假设 maxLen 为原始的最长长度)
	// ab => ^#a#b#$ ; maxRadius = 1, centerPos = 2 ;转换之后 maxLen = 1, centerPos = 0
	// aa => ^#a#a#$ ; maxRadius = 2, centerPos = 3 ;转换之后 maxLen = 2, centerPos = 0
	// aba => ^#a#b#a#$ ; maxRadius = 3, centerPos = 4 ;转换之后 maxLen = 3, centerPos = 1
	centerPos = centerPos/2 - 1

	// 根据中心点计算起始位置，重用下 centerPos
	centerPos = centerPos - (maxRadius-1)/2

	return s[centerPos : centerPos+maxRadius]
}
