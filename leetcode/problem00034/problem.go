package problem00034

import (
	"sort"
)

func searchRange(nums []int, target int) []int {
	length := len(nums)

	switch length {
	case 0:
		return []int{-1, -1}
	case 1:
		if target == nums[0] {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	// 如果小于最小值，大于最大值，则肯定不存在
	if nums[0] > target || nums[length-1] < target {
		return []int{-1, -1}
	}

	// 找到第一个满足，nums[first] >= target
	first := sort.SearchInts(nums, target)
	// 如果没找到
	if length == first || target != nums[first] {
		return []int{-1, -1}
	}
	// 此处找到了一个
	// 如果只有一个数是 target 的话
	if first == length-1 || nums[first+1] > target {
		return []int{first, first}
	}

	// 在 first + 1 后面找到第一个满足，nums[second] >= target + 1
	second := sort.SearchInts(nums[first+1:], target+1) + first + 1

	return []int{first, second - 1}
}
