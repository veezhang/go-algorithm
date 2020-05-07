package problem00028

import (
	"fmt"
	"strings"
)

// 时间复杂度：	O((T - P)*P)
// 空间复杂度：	O(1)
func strStrSubstring(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// 主串 T(Text) ，下标 i ， 截取长度 lenP : T[i:i+lenP]
	// 剩余长度需要 >= lenP , lenT - i >= lenP , i <= lenT - lenP
	for i := 0; i <= lenT-lenP; i++ {
		if T[i:i+lenP] == P {
			return i
		}
	}

	return -1
}

// 时间复杂度：	O((T - P)*P)
// 空间复杂度：	O(1)
func strStrTwoPointer(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// 主串 T(Text) ，下标 i(i 为某一轮中主串 S 的起始位置，比较使用 S[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。
	// 剩余长度需要 >= lenP , lenT - i >= lenP , i <= lenT - lenP
	for i := 0; i <= lenT-lenP; i++ {
		// 依次与模式串 P 比较
		// j 最大为 lenP - 1 ； i+j 最大为： i + lenP - 1 ，最大为： lenT - lenP + lenP - 1 = lenT - 1 ， 因此不会越界
		j := 0
		for T[i+j] == P[j] {
			j++

			if j == lenP {
				return i
			}
		}
	}

	return -1
}

// 时间复杂度：	O(T + P)
// 空间复杂度：	O(1)
func strStrKMP(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// next 是最长前后缀数组右移一位
	next := make([]int, lenP)
	// 计算最长公共前后缀长度
	next[0] = -1
	for i, j := 0, -1; i < lenP-1; {
		// i 表示最长公共前后缀中后缀的下标
		// j 表示最长公共前后缀中前缀的下标，也正好是前一个最长公共前后缀的长度，也正好也就是前缀中最后一个字符的后一个字符的下标
		if j == -1 || P[i] == P[j] {
			i++
			j++
			// 当后面进行匹配的时候，如果 T[i] != P[j] ， 则会比较 T[i] ?= P[next[j]] ，
			// 此时，如果 P[j] = P[ next[j] ] ，则 T[i] != P[next[j]] ，会继续匹配失败
			// 所以 P[j] = P[ next[j] ] 的时候优化下， next[j] = next[ next[j] ] ，继续往前找一个不相等的
			//
			// 如果 P[i] == P[j] 的时候设置为 next[i] = j
			// 则正好 P[j] = P[ next[j] ] 这种情况会出现
			if P[i] == P[j] {
				// 因为不能出现p[i] = p[ next[i]]，所以当出现时需要继续递归，k = next[k] = next[next[k]]
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			// 如果不相等，则比较最长前缀的最长前缀
			j = next[j]
		}
	}

	// i 为主串中的下标， j 为匹配串中的下标
	i, j := 0, 0
	for i < lenT && j < lenP {
		// 如果 j = -1 ，或者当前字符匹配成功（即 T[i] == P[j] ），则 i++,j++
		if j == -1 || T[i] == P[j] {
			i++
			j++
		} else {
			//如果 j != -1 ，且当前字符匹配失败（即 T[i] == P[j] ），则令 i 不变，j = next[j]
			j = next[j]
		}
	}

	if j == lenP {
		return i - j
	}

	return -1
}

// 时间复杂度：	最好： O(T/P) 最差： O(T*P)
// 空间复杂度：	O(S) S是与P, T相关的有限字符集长度，我们这里直接用的 256
func strStrBM(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// 坏字符
	// shiftBadTable 为模式串 P 中坏字符规则需要移动的位数
	// 不存在的字符，移动位数 = lenP ， 否则移动位数 = lenP - i - 1
	shiftBadTable := [256]int{}
	for i := 0; i < 256; i++ {
		shiftBadTable[i] = lenP
	}
	for i := 0; i < lenP; i++ {
		shiftBadTable[P[i]] = lenP - i - 1
	}

	// 好后缀
	// shiftGoodTable[i] 表示遇到好后缀时，模式串应该移动的距离，其中 i 表示好后缀前面一个字符的位置（也就是坏字符的位置）
	shiftGoodTable := make([]int, lenP)
	// suffix[i] = l => 表示以i为边界，与模式串后缀匹配的最大长度，满足 P[i-l, i] == P[lenP-1-l, lenP-1]
	suffix := make([]int, lenP)
	suffix[lenP-1] = lenP
	for i := lenP - 2; i >= 0; i-- {
		l := 0
		for i-l >= 0 && P[i-l] == P[lenP-1-l] {
			l++
		}
		suffix[i] = l
	}

	// 好后缀有三种情况
	// Case1: 模式串中存在好后辍构成的子串，这种情况只需要将模式串中最靠右的子串与好后辍对齐即可。
	// Case2: 模式串中没有子串匹配上好后缀，但是存在与好后辍匹配的前辍，此时只需要将好后辍的后辍与模式串中对应的最长前辍匹配即可。
	// Case3: 模式串中没有子串匹配上好后缀，也不存在与好后辍匹配的前辍，则需要将模式串整个移动到目标串中不匹配位置之后即可。

	// 3种情况获得的 shiftGoodTable[i] 值比较 ： 3 > 2 > 1
	// 为了保证其值越来越小，所以按顺序处理 3->2->1 情况

	// Case3: 全部初始为模式串 P 的长度
	for i := 0; i < lenP; i++ {
		shiftGoodTable[i] = lenP
	}

	// 处理第二种情况
	// 需要 shiftGoodTable[j] = lenP - 1 - i 取值是小的，所以 i 从大到小，则 lenP - 1 - i 从小到大，只会设置一次，设置的就是小的
	for i := lenP - 1; i >= 0; i-- {
		// suffix[i] == i+1 表示 i 及之前的等于模式串 P 的后面的
		// 也就是模式串的前缀 = 后缀，满足 Case2
		if suffix[i] == i+1 { // 找到合适位置
			// lenP - 1 - i 为匹配到的后缀的第一个下标，同时也是 i 到末尾的距离
			// 坏字符是 [0,lenP - 1 - i) 区间中的时候，都设置为 lenP - 1 - i ，也就是出现怀着字符的时候，移动 lenP - 1 - i ，让前缀和后缀对齐
			//
			// 例如： P = BCDABC ， lenP = 6 ， i = 1 ， suffix[i] = 2 ，lenP-1-i = 6 - 1 - 1 = 4
			// BCDA 对应的是坏字符的话，都让 BC 和 BC 对齐
			for j := 0; j < lenP-1-i; j++ {
				if shiftGoodTable[j] == lenP { // 保证每个位置至多只能被修改一次
					shiftGoodTable[j] = lenP - 1 - i
				}
			}
		}
	}

	// 1.处理第一种情况，顺序是从前到后
	// 需要 shiftGoodTable[lenP-1-suffix[i]] = lenP - 1 - i 取值是小的，所以 i 从小到大，则 lenP - 1 - i 从大到小，后面的设置覆盖前面的，设置的就是小的
	for i := 0; i < lenP-1; i++ {
		// suffix[i] => P[i-l, i] == P[lenP-1-l, lenP-1]
		// suffix[i] 为后缀的长度，lenP-1-suffix[i] 为好后缀前一个字符，也就是坏字符， lenP - 1 - i 是 i 到末尾的距离，也就是移动的距离。
		//
		// 假设前后有两个 i1，i2 (i1 < i2) 都的 suffix[i] 值一样，则 lenP - 1 - i1 > lenP - 1 - i2 ，最终会取 lenP - 1 - i2 ，也就是移动相对较短的。
		// P = BCABCDABCEABC, lenP = 13, i1 = 4, i2 = 8, suffix[i1] = suffix[i2] = 3
		// shiftGoodTable[13 - 1 - 3] = shiftGoodTable[9] = 13 - 1 - 4 = 8
		// shiftGoodTable[13 - 1 - 8] = shiftGoodTable[9] = 13 - 1 - 8 = 4
		shiftGoodTable[lenP-1-suffix[i]] = lenP - 1 - i
	}

	// 开始 BM
	for i := 0; i <= lenT-lenP; {
		// 依次与模式串 P 比较
		// j 最大为 lenP - 1 ； i+j 最大为： i + lenP - 1 ，最大为： lenT - lenP + lenP - 1 = lenT - 1 ， 因此不会越界
		j := lenP - 1
		for T[i+j] == P[j] {
			j--
			if j < 0 {
				return i
			}
		}
		// 此处 T[i+j] != P[j] , T[i+j] 为坏字符, T[i+j+1:i+j+1+lenP-1-j] 为好后缀，即 T[i+j+1:i+lenP]

		// 根据坏字符规则移动
		// 坏字符都结尾长度：shiftBadTable[T[i+j]] ， j 到结尾的长度： ((lenP - 1) - j)
		// shiftBadTable[T[i+j]] - ((lenP - 1) - j) 就表示坏字符在模式串 P 中的位值跟 j 对齐
		move := shiftBadTable[T[i+j]] - ((lenP - 1) - j)
		if shiftGoodTable[j] > move {
			// 根据好后缀规则移动
			move = shiftGoodTable[j]
		}

		if false {

			fmt.Printf("\ni = %d j = %d move = (%d %d %d), i+move = %d\n", i, j, shiftBadTable[T[i+j]]-(lenP-1)+j, shiftGoodTable[j], move, i+move)
			// 打印 i
			fmt.Printf("i     ")
			for ii := 0; ii < lenT; ii++ {
				fmt.Printf("%3d", ii)
			}
			fmt.Println()
			// 打印 T
			fmt.Printf("T     ")
			for ii := 0; ii < lenT; ii++ {
				fmt.Printf("%3c", T[ii])
			}
			fmt.Println()
			// 打印 √ ×
			fmt.Printf("      ")
			for ii := 0; ii < i+j; ii++ {
				fmt.Printf("%3s", "")
			}
			fmt.Printf("%3c", '×')
			for ii := 0; ii < lenP-j-1; ii++ {
				fmt.Printf("%3c", '√')
			}
			fmt.Println()
			// 打印 j
			fmt.Printf("j     ")
			for ii := 0; ii < i; ii++ {
				fmt.Printf("%3s", "")
			}
			for ii := 0; ii < lenP; ii++ {
				fmt.Printf("%3d", ii)
			}
			fmt.Println()
			// 打印 P
			fmt.Printf("P     ")
			for ii := 0; ii < i; ii++ {
				fmt.Printf("%3s", "")
			}
			for ii := 0; ii < lenP; ii++ {
				fmt.Printf("%3c", P[ii])
			}
			fmt.Println()
			// 打印 Bad
			fmt.Printf("Bad   ")
			for ii := 0; ii < i; ii++ {
				fmt.Printf("%3s", "")
			}
			for ii := 0; ii < lenP; ii++ {
				fmt.Printf("%3d", shiftBadTable[P[ii]])
			}
			fmt.Println()
			// 打印 suffix
			fmt.Printf("suffix")
			for ii := 0; ii < i; ii++ {
				fmt.Printf("%3s", "")
			}
			for ii := 0; ii < lenP; ii++ {
				fmt.Printf("%3d", suffix[ii])
			}
			fmt.Println()
			// 打印 Good
			fmt.Printf("Good  ")
			for ii := 0; ii < i; ii++ {
				fmt.Printf("%3s", "")
			}
			for ii := 0; ii < lenP; ii++ {
				fmt.Printf("%3d", shiftGoodTable[ii])
			}
			fmt.Println()
		}

		i += move
	}
	return -1
}

func strStrHorspool(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// shiftTable 为模式串 P 中字符需要移动的位数
	// 不存在的字符，移动位数 = lenP ， 否则移动位数 = lenP - i - 1
	shiftTable := [256]int{}
	for i := 0; i < 256; i++ {
		shiftTable[i] = lenP
	}
	// 模式串中最后一个不处理，因为我们取的是最后一个对应的字符，如果 P[lenP-1] 正要与之相等，那么每次拿这个与之对应，移动 0 位
	for i := 0; i < lenP-1; i++ {
		shiftTable[P[i]] = lenP - i - 1
	}

	// 主串 T(Text) ，下标 i(i 为某一轮中主串 S 的起始位置，比较使用 S[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。
	// 剩余长度需要 >= lenP , lenT - i >= lenP , i <= lenT - lenP
	for i := 0; i <= lenT-lenP; {
		// 依次与模式串 P 比较
		// j 最大为 lenP - 1 ； i+j 最大为： i + lenP - 1 ，最大为： lenT - lenP + lenP - 1 = lenT - 1 ， 因此不会越界
		j := lenP - 1
		for T[i+j] == P[j] {
			j--
			if j < 0 {
				return i
			}
		}
		// 此处 T[i+j] != P[j]

		// 计算移动位置
		i += shiftTable[T[i+lenP-1]]
	}

	return -1
}

// 时间复杂度：	O(T) 最差： O(T*P)
// 空间复杂度：	O(1)
func strStrSunday(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// 主串 T(Text) ，下标 i(i 为某一轮中主串 S 的起始位置，比较使用 S[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。
	// 剩余长度需要 >= lenP , lenT - i >= lenP , i <= lenT - lenP
	for i := 0; i <= lenT-lenP; {
		// 依次与模式串 P 比较
		// j 最大为 lenP - 1 ； i+j 最大为： i + lenP - 1 ，最大为： lenT - lenP + lenP - 1 = lenT - 1 ， 因此不会越界
		j := 0
		for T[i+j] == P[j] {
			j++

			if j == lenP {
				return i
			}
		}
		// 取下一个字符 i + lenP
		// i = lenT-lenP 时， i + lenP = lenT ， 越界
		if i+lenP >= lenT {
			return -1
		}
		// 计算移动位置
		for j = lenP - 1; j >= 0; j-- {
			if T[i+lenP] == P[j] {
				break
			}
		}
		// 如果没有找到 j = -1
		// if j == -1 {
		// 	i += lenP + 1
		// } else {
		// 	i += lenP - j
		// }
		i += lenP - j
	}

	return -1
}

// 时间复杂度：	O(T) 最差： O(T*P)
// 空间复杂度：	O(1)
func strStrSundayWithShiftTable(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// shiftTable 为模式串 P 中字符需要移动的位数
	// 不存在的字符，移动位数 = lenP + 1 ， 否则移动位数 = lenP - i
	shiftTable := [256]int{}
	for i := 0; i < 256; i++ {
		shiftTable[i] = lenP + 1
	}
	for i := 0; i < lenP; i++ {
		shiftTable[P[i]] = lenP - i
	}

	// 主串 T(Text) ，下标 i(i 为某一轮中主串 S 的起始位置，比较使用 S[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。
	// 剩余长度需要 >= lenP , lenT - i >= lenP , i <= lenT - lenP
	for i := 0; i <= lenT-lenP; {
		// 依次与模式串 P 比较
		// j 最大为 lenP - 1 ； i+j 最大为： i + lenP - 1 ，最大为： lenT - lenP + lenP - 1 = lenT - 1 ， 因此不会越界
		j := 0
		for T[i+j] == P[j] {
			j++

			if j == lenP {
				return i
			}
		}
		// 取下一个字符 i + lenP
		// i = lenT-lenP 时， i + lenP = lenT ， 越界
		if i+lenP >= lenT {
			return -1
		}
		// 计算移动位置
		i += shiftTable[T[i+lenP]]
	}

	return -1
}

// 时间复杂度：	O(T + P)
// 空间复杂度：	O(1)
func strStrRabinKarp(T string, P string) int {
	lenT, lenP := len(T), len(P)
	if lenP == 0 {
		return 0
	}
	if lenT < lenP {
		return -1
	}

	// primeRK is the prime base used in Rabin-Karp algorithm.
	const primeRK = 16777619

	// hashStr returns the hash and the appropriate multiplicative
	// factor for use in Rabin-Karp algorithm.
	hashStr := func(sep string) (uint32, uint32) {
		hash := uint32(0)
		for i := 0; i < len(sep); i++ {
			hash = hash*primeRK + uint32(sep[i])
		}
		var pow, sq uint32 = 1, primeRK
		for i := len(sep); i > 0; i >>= 1 {
			if i&1 != 0 {
				pow *= sq
			}
			// 只有32位，超出范围的会被丢掉
			sq *= sq
		}
		return hash, pow
	}

	hashss, pow := hashStr(P)
	var h uint32
	for i := 0; i < lenP; i++ {
		h = h*primeRK + uint32(T[i])
	}
	if h == hashss && T[:lenP] == P {
		return 0
	}
	for i := lenP; i < lenT; {
		h *= primeRK
		h += uint32(T[i])
		h -= pow * uint32(T[i-lenP])
		i++
		if h == hashss && T[i-lenP:i] == P {
			return i - lenP
		}
	}
	return -1
}

func strStrGO(T string, P string) int {
	return strings.Index(T, P)
}
