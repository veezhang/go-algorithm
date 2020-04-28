package problem00016

import "sort"

// 双指针
func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	// 最小的4个数和大于 target 或者最大的4个数和小于 target ，应为是递增的，结果4个数肯定要小于或大于 target
	if nums[0]+nums[1]+nums[2]+nums[3] > target || nums[length-4]+nums[length-3]+nums[length-2]+nums[length-1] < target {
		return [][]int{}
	}

	ret := make([][]int, 0)

	// 取第一个数字，然后双指针获取另外两个数字， i < length-2
	for i := 0; i < length-3; {
		for j := i + 1; j < length-2; {
			left, right := j+1, length-1
			for left < right {
				// 最小的4个数和大于 target 或者最大的4个数和小于 target ，应为是递增的，结果4个数肯定要小于或大于 target
				// i < length-3，所以 i+3 < length 没有越界
				// right > left > j > i 所以 right 也不会越界
				if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target || nums[right-3]+nums[right-2]+nums[right-1]+nums[right] < target {
					break
				}

				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
				}

				// 注意： 等于 target 的情况也是需要处理的。
				// 比如 2 = -7 -2 2 9，后面还是有可能有 -7 -2 3 8 的

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
				} else { // 等于 0 的情况，左右一起移动，因为不会出现重复的， 那么 left 和 right 都不会出现了
					// left 右移， 过滤重复值， left < right ，所以 left+1 最大等于 right 不会越界
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
					// right 左移， 过滤重复值， right > left ，所以 right-1 最小等于 left 不会越界
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				}
			}
			// 过滤重复的 j
			for j < length-2 && nums[j] == nums[j+1] {
				j++
			}
			j++
		}
		// 过滤重复的 i
		for i < length-3 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return ret
}
