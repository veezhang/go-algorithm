package problem00030

func findSubstringLoop(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	ret := make([]int, 0)

	wordCount := len(words)
	wordLength := len(words[0])
	totalLength := wordCount * wordLength

	wordsMap := make(map[string]int, wordCount)

	for i := 0; i < wordCount; i++ {
		wordsMap[words[i]] = wordsMap[words[i]] + 1
	}

	for i := 0; i < len(s)-totalLength+1; i++ {
		currentStr := s[i : i+wordLength]
		findCount := 0
		var tmpMap map[string]int

		for n, ok := wordsMap[currentStr]; ok; {
			if tmpMap == nil {
				tmpMap = make(map[string]int, wordCount)
			}
			nn := tmpMap[currentStr]
			if nn >= n {
				break
			}
			tmpMap[currentStr] = nn + 1
			findCount++

			//  如果找到了，添加到结果中
			if findCount == wordCount {
				ret = append(ret, i)
				break
			}

			currentStr = s[i+findCount*wordLength : i+findCount*wordLength+wordLength]
			n, ok = wordsMap[currentStr]
		}
	}

	return ret
}

func findSubstringSlidingWindow(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	ret := make([]int, 0)

	wordCount := len(words)
	wordLength := len(words[0])
	totalLength := wordCount * wordLength

	wordsMap := make(map[string]int, wordCount)

	for i := 0; i < wordCount; i++ {
		wordsMap[words[i]] = wordsMap[words[i]] + 1
	}

	for i := 0; i < wordLength; i++ {
		// 每一个区间在  [left, right) 中，不包含 right
		// right = left + totalLength
		left := i
		findCount := 0
		var tmpMap map[string]int

		// 剩余长度够 totalLength
		for left+totalLength <= len(s) {
			// 以下 for 循环比较最后一个是否包含在 words 中
			bLastContains := false
			// 找匹配到的数目
			for findCount < wordCount {
				currentStr := s[left+findCount*wordLength : left+findCount*wordLength+wordLength]
				n, ok := wordsMap[currentStr]
				if !ok {
					break
				}
				bLastContains = true
				if tmpMap == nil {
					tmpMap = make(map[string]int, wordCount)
				}
				nn := tmpMap[currentStr]
				if nn >= n {
					break
				}
				tmpMap[currentStr] = nn + 1
				findCount++
			}

			//  如果找到了，添加到结果中
			if findCount == wordCount {
				ret = append(ret, left)
			}

			// 如果不满足条件的时候， words 包含，这里移动 left ，并且取消最左边对 tmpMap 的影响
			if bLastContains {
				// 移除左边的一个字符，并减少对应的计数器
				tmpMap[s[left:left+wordLength]] = tmpMap[s[left:left+wordLength]] - 1
				findCount--

				// 移动窗口
				left += wordLength
			} else {
				// 如果 words 不包含，则移动到最后一个未匹配后面开始找
				// 因为 left 到 最后一个未匹配 不可能中得到了
				// 注意这里不能仅仅是根据是否找到了 findCount 个，因为可能是由于相同的字符多了导致没有找到 findCount 个匹配的
				left += (findCount + 1) * wordLength
				findCount = 0
				tmpMap = make(map[string]int, wordCount)
			}
		}
	}

	return ret
}
