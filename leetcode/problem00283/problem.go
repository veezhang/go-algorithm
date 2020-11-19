package problem00283

func moveZeroes(nums []int) {
	// left 已经非零的尾部（下一个非零放到这里）
	// right 待处理数据
	for left, right, n := 0, 0, len(nums); right < n; right++ {
		// 如果不为零，则换下，并且 left++
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
