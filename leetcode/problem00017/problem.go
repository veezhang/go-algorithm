package problem00017

import (
	"unsafe"
)

func letterCombinationsDirect(digits string) []string {
	length := len(digits)
	if length == 0 {
		return []string{}
	}

	lettersTable := [][]byte{
		{},                   // 0
		{},                   // 1
		{'a', 'b', 'c'},      // 2
		{'d', 'e', 'f'},      // 3
		{'g', 'h', 'i'},      // 4
		{'j', 'k', 'l'},      // 5
		{'m', 'n', 'o'},      // 6
		{'p', 'q', 'r', 's'}, // 7
		{'t', 'u', 'v'},      // 8
		{'w', 'x', 'y', 'z'}, // 9
	}

	// total 为总结果有多少个可能的字符串
	// strLen 为每一个字符串的长度
	// 统计总共有多少种结果，以及每种结果字符串的长度
	// 因为仅包含 2-9 的数字，所以也不用考虑 len(lettersTable[digits[i]-'0']) = 0 的情况，不然需要过滤下 digits
	total := 0
	strLen := 0
	for i := 0; i < length; i++ {
		if total == 0 {
			total = 1
		}
		total *= len(lettersTable[digits[i]-'0'])
		strLen++
	}

	// ret 会需要返回的
	// currRetLen 为 ret 当前已经设置过的长度
	// 每一个位置，是什么字符基本上都是固定的，后面只需要用索引就可以了
	// 减少不必要的 append 和扩充相关的
	ret := make([][]byte, total, total)
	for i := 0; i < total; i++ {
		ret[i] = make([]byte, strLen, strLen)
	}
	currRetLen := 0

	// 遍历每一个数字
	for i := 0; i < length; i++ {
		// 获取数字对应的字母
		letters := lettersTable[digits[i]-'0']
		lettersLen := len(letters)

		// 如果没有字符，继续找下一个
		if lettersLen == 0 {
			continue
		}

		// 假设当前有 currRetLen 而下一个字符对应有 lettersLen 个字符
		// 那么拿 ret 中每个分别加上 letters 中的字符，则可以生成 currRetLen * lettersLen 个可能
		// 由于遍历第一个数字的时候， currRetLen = 0 ，这种情况直接等于 lettersLen

		// 由于赋值依赖 [0, currRetLen) 区间中的值，所以这里从后往前进行赋值，避免把之前的数据覆盖了
		// 比如下面的， 如果先顺着来，修改 ret[1] 的时候依赖 ret[0] ，但是 ret[1] 已经修改了

		// 遍历 lettersLen 的时候，我们也倒着来好了，让字符能够升序，根据题意，顺着也无妨
		// 发现 letters 可以直接用取模来进行设置，不再需要单独遍历

		// 比如： 23
		// i = 0 的时候, letters = {'a', 'b', 'c'}, currRetLen = 0, lettersLen = 3, nextRetLen = 3
		// ret[2] = {'c'}
		// ret[1] = {'b'}
		// ret[0] = {'a'}
		//
		// i = 1 的时候, letters = {'d', 'e', 'f'}, currRetLen = 3, lettersLen = 3, nextRetLen = 9
		// ret[8] = ret[8/3] + letters[8%3] = ret[2] + letters[2] = {'c'} + {'f'} = {'c','f'}
		// ret[7] = ret[7/3] + letters[7%3] = ret[2] + letters[1] = {'c'} + {'e'} = {'c','e'}
		// ret[6] = ret[6/3] + letters[6%3] = ret[2] + letters[0] = {'c'} + {'d'} = {'c','d'}
		// ret[5] = ret[5/3] + letters[5%3] = ret[1] + letters[2] = {'b'} + {'f'} = {'b','f'}
		// ret[4] = ret[4/3] + letters[4%3] = ret[1] + letters[1] = {'b'} + {'e'} = {'b','e'}
		// ret[3] = ret[3/3] + letters[3%3] = ret[1] + letters[0] = {'b'} + {'d'} = {'b','d'}
		// ret[2] = ret[2/3] + letters[2%3] = ret[0] + letters[2] = {'a'} + {'f'} = {'a','f'}
		// ret[1] = ret[1/3] + letters[1%3] = ret[0] + letters[1] = {'a'} + {'e'} = {'a','e'}
		// ret[0] = ret[0/3] + letters[0%3] = ret[0] + letters[0] = {'a'} + {'d'} = {'a','d'}

		nextRetLen := currRetLen * lettersLen
		if nextRetLen == 0 {
			nextRetLen = lettersLen
		}

		// 遍历接下来的所有可能情况进行赋值
		// 这里从后往前进行赋值，避免把之前的数据覆盖了
		for j := nextRetLen - 1; j >= 0; j-- {
			// 将之前的 i-1 个复制过来, 也就是 ret[j/lettersLen][:i] 。后面也就只需要设置当前第 i 个字符了
			copy(ret[j][:i], ret[j/lettersLen][:i])

			// 遍历将每一个字符追加到之前的， 我们也倒着来好了，让字符能够升序，根据题意，顺着也无妨
			ret[j][i] = letters[j%lettersLen]
		}

		currRetLen = nextRetLen
	}

	retStr := make([]string, total, total)
	for i := 0; i < total; i++ {
		// byte to string
		retStr[i] = *(*string)(unsafe.Pointer(&ret[i]))
	}

	return retStr
}

var lettersTable = [][]byte{
	{},                   // 0
	{},                   // 1
	{'a', 'b', 'c'},      // 2
	{'d', 'e', 'f'},      // 3
	{'g', 'h', 'i'},      // 4
	{'j', 'k', 'l'},      // 5
	{'m', 'n', 'o'},      // 6
	{'p', 'q', 'r', 's'}, // 7
	{'t', 'u', 'v'},      // 8
	{'w', 'x', 'y', 'z'}, // 9
}

func letterCombinationsBacktrackOnce(combinations string, digits string, index int, result *[]string) {
	// 满足结束条件
	if index >= len(digits) {
		*result = append(*result, combinations)
		return
	}

	// 获取数字对应的字母
	letters := lettersTable[digits[index]-'0']
	// for 选择 in 选择列表:
	for i := 0; i < len(letters); i++ {
		// 将对应的字符添加到 combinations 中，并继续处理后面的
		// 做选择 combinations = combinations+string([]byte{letters[i]})
		letterCombinationsBacktrackOnce(combinations+string([]byte{letters[i]}), digits, index+1, result)
		// 撤销选择 combinations 还是之前的
	}
}

func letterCombinationsBacktrack(digits string) []string {
	length := len(digits)
	if length == 0 {
		return []string{}
	}

	ret := make([]string, 0)

	// ret append 的时候会修改 ret ， 所以这里使用 &ret
	letterCombinationsBacktrackOnce("", digits, 0, &ret)

	return ret
}
