package problem00321

// 从数组中选出 k 个数组成新的数组，使其为最大数，需保持相对顺序。
func maxNumberOne(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}
	length := len(nums)
	if k >= length {
		return nums
	}
	// 需要丢弃元素个数
	drop := length - k
	stack := make([]int, 0, k)
	for _, v := range nums {
		// 如果栈顶元素 < v 则移除栈顶
		for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < v {
			// for k > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
			drop--
		}

		// 如果元素个数还不到 k 个，则继续添加
		if len(stack) < k {
			stack = append(stack, v)
		} else {
			// 否则算丢弃
			drop--
			if 0 == drop {
				// 如果都丢弃完了，则跳出循环
				break
			}
		}
	}
	return stack
}

// 合并两个数，使其为最大数，需保持相对顺序。
func mergeNumber(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	if 0 == l1 {
		return nums2
	}

	if 0 == l2 {
		return nums1
	}

	merged := make([]int, l1+l2)

	for i := range merged {
		if less(nums1, nums2) {
			merged[i], nums2 = nums2[0], nums2[1:]
		} else {
			merged[i], nums1 = nums1[0], nums1[1:]
		}
	}

	return merged
}

// 判断是否大于
func less(nums1 []int, nums2 []int) bool {
	l1, l2 := len(nums1), len(nums2)
	for i := 0; i < l1 && i < l2; i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return l1 < l2
}

func maxNumber(nums1 []int, nums2 []int, k int) (ans []int) {
	if k <= 0 {
		return
	}
	l1, l2 := len(nums1), len(nums2)
	if 0 == l1 {
		return maxNumberOne(nums2, k)
	}

	if 0 == l2 {
		return maxNumberOne(nums1, k)
	}

	// 保证 nums1 为较短的数组
	if l1 > l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}
	// k1, k2 分别为 nums1, nums2 的获取的长度
	// k1 <= l1, k1 <= k
	// k2 := k - k1 <= l2 ==> k1 >= k - l2
	k1 := 0
	if k > l2 {
		k1 = k - l2
	}
	for ; k1 <= k && k1 <= l1; k1++ {
		nums1k := maxNumberOne(nums1, k1)
		nums2k := maxNumberOne(nums2, k-k1)
		merged := mergeNumber(nums1k, nums2k)

		if less(ans, merged) {
			ans = merged
		}

	}

	return
}
