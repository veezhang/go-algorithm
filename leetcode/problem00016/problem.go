package problem00016

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 双指针
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		panic("length < 3")
	}

	sort.Ints(nums)

	ret := nums[0] + nums[1] + nums[2]

	// 取第一个数字，然后双指针获取另外两个数字， i < length-2
	for i := 0; i < length-2; {
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(ret-target) {
				ret = sum
			}

			if sum < target { // 数字小了， left 往右移动
				// left 右移， 过滤重复值， left < right ，所以 left+1 最大等于 right 不会越界
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if sum > target { // 数字大了， left 往右移动
				// right 左移， 过滤重复值， right > left ，所以 right-1 最小等于 left 不会越界
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else { // 等于 target 的情况，可以返回了
				return ret
			}
		}
		// 过滤重复的 i
		for i < length-2 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return ret
}
