package problem00026

// 时间复杂度： O(n)
// 空间复杂度： O(1)
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// index 为当前不重复的数下标，第一个数字已经不重复了， i 从 1 开始
	// 下一个不重复的将放置到 nums[index+1] ，最后长度为 index+1
	index := 0
	for i := 1; i < length; i++ {
		if nums[i] != nums[index] {
			index++
			// index 肯定小于 i ， 这里不判断 index == i 的情况
			nums[index] = nums[i]
		}
	}

	return index + 1
}
