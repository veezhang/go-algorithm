package problem00015

import "sort"

func threeSumHash(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	// 第二级 map 用第三个数 num3 作为 key ， 第三个数字 num2 的索引作为 value ， 使用过后索引置为 -1
	// 第一级 map 用第一个数 num1 作为 key
	m := make(map[int]map[int]int, length)
	ret := make([][]int, 0)

	for i := 0; i < length; i++ {
		num1 := nums[i]
		mm, ok := m[num1]
		if !ok {
			mm = make(map[int]int)
			m[num1] = mm
		}
		for j := i + 1; j < length; j++ {
			num2 := nums[j]
			num3 := 0 - num1 - num2

			// 如果 mm 中存在第三个数字，则判断是否统计过了，没有统计过则添加到 ret 作为答案，并且标记 num2 索引为 -1；否则继续
			// 如果 mm 没有，则设置下 num2 的索引
			idx, ok := mm[num3]
			if ok {
				// id < 0 表示已经统计过了
				// idx == i || idx == j 表示是重复的，这里 i 已经不等于 j 了
				if idx < 0 || idx == i || idx == j {
					continue
				}

				// 判断三个数 num1, num2, num3 是否已经使用过
				//
				// num1, num2, num3
				// num1, num3, num2
				// num2, num1, num3
				// num2, num3, num1
				// num3, num1, num2
				// num3, num2, num1

				if mm1, ok1 := m[num1]; ok1 {
					if idx, ok2 := mm1[num2]; ok2 && idx < 0 {
						continue
					}

					if idx, ok2 := mm1[num3]; ok2 && idx < 0 {
						continue
					}
				}

				if mm1, ok1 := m[num2]; ok1 {
					if idx, ok2 := mm1[num1]; ok2 && idx < 0 {
						continue
					}

					if idx, ok2 := mm1[num3]; ok2 && idx < 0 {
						continue
					}
				}

				if mm1, ok1 := m[num3]; ok1 {
					if idx, ok2 := mm1[num1]; ok2 && idx < 0 {
						continue
					}

					if idx, ok2 := mm1[num2]; ok2 && idx < 0 {
						continue
					}
				}

				ret = append(ret, []int{num1, num2, num3})
				mm[num2] = -1
			} else {
				mm[num2] = j
			}
		}
	}

	return ret
}

// 双指针
func threeSumTwoPointer(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	// 全是 0
	if nums[0] == 0 && nums[length-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	// 全为正数或者全为负数，应为是递增的，只用判断边界
	if nums[0] > 0 || nums[length-1] < 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)

	// 取第一个数字，然后双指针获取另外两个数字， i < length-2
	for i := 0; i < length-2; {
		left, right := i+1, length-1
		for left < right {
			// 全为正数或者全为负数，应为是递增的，只用判断边界
			if nums[i] > 0 || nums[right] < 0 {
				break
			}

			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
			}

			// 注意： 等于 0 的情况也是需要处理的。
			// 比如 -7 -2 9，后面还是有可能有 -7 2 5 的

			if sum < 0 { // 数字小了， left 往右移动
				// left 右移， 过滤重复值， left < right ，所以 left+1 最大等于 right 不会越界
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if sum > 0 { // 数字大了， left 往右移动
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
		// 过滤重复的 i
		for i < length-2 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return ret
}
