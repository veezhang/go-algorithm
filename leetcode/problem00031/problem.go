package problem00031

func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	// 1. 先找出最大的索引 k 满足 nums[k] < nums[k+1]，则表示 [k+1, n) 为降序的； [k, n) 为一个倒对勾形状
	k := length - 2
	for ; k >= 0 && nums[k] >= nums[k+1]; k-- {
	}
	if k >= 0 {
		// 2. 再找出另一个最大索引 l (l > k) 满足 nums[l] > nums[k]；肯定能找到，因为 nums[k+1]> nums[k]
		l := length - 1
		for ; nums[l] <= nums[k]; l-- { // 这里肯定能找到，就不判断 l >= 0 了
		}
		// 3. 交换 nums[l] 和 nums[k]；此时 [k+1, n) 一定还是降序，因为 nums[l] 是从右侧起第一个大于 nums[k] 的值。
		nums[l], nums[k] = nums[k], nums[l]
	}
	// 5. 如果 1 中不存在，则整个为降序，已经是最大排列，翻转整个数组得到最小排列；
	// 4. 反转 nums[k+1:] ，变成升序。
	reverse(nums[k+1:])
}

func reverse(a []int) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}
