package problem00290

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// word2ch s 中字符串到 pattern 中字符的映射
	// word2ch pattern 中字符到 s 中字符串的映射
	word2ch := make(map[string]byte, len(pattern))
	ch2word := make(map[byte]string, len(pattern))

	for i, word := range words {
		ch := pattern[i]
		// 如果 s 中字符串出现过，但是和 pattern 中字符不对应，返回 false
		if c, ok := word2ch[word]; ok && ch != c {
			return false
		}
		// 如果 pattern 中字符出现过，但是和 s 中字符串不对应，返回 false
		if w, ok := ch2word[ch]; ok && word != w {
			return false
		}
		// 统计互相映射关系
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
