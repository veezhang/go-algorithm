package problem00842

import "sort"

func groupAnagrams1(strs []string) [][]string {
	// 用 map 统计异位词，异位词则每个排序后字符一致
	mp := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		sortedStr := string(bytes)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams2(strs []string) [][]string {
	// 用 map 统计异位词，异位词则每个字符个数相同
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
