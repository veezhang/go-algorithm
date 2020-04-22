package problem00004

// K值法
func findMedianSortedArraysKth(nums1 []int, nums2 []int) float64 {
	length1, length2 := len(nums1), len(nums2)

	total_length := length1 + length2

	// index1,index2 分别为nums1，nums2中的下标，过滤较小的时候，对应的下标往后移动
	index1, index2 := 0, 0

	// 找第 (total_length + 1) / 2 个数
	k := (total_length + 1) / 2

	// 循环到k=1(找最小的那个数字)，或者有一个已经不够过滤了
	for {
		// nums1 已经遍历完成了，直接在nums2中取
		if index1 >= length1 {
			if total_length%2 == 0 {
				return ((float64)(nums2[index2+k-1] + nums2[index2+k])) / 2
			}
			return (float64)(nums2[index2+k-1])
		}

		// nums2 已经遍历完成了，直接在nums1中取
		if index2 >= length2 {
			if total_length%2 == 0 {
				return ((float64)(nums1[index1+k-1] + nums1[index1+k])) / 2
			}
			return (float64)(nums1[index1+k-1])
		}

		// 找到最后一个了
		if k == 1 {
			// 求一个小的，得到第一个数字
			m1 := 0
			mIn1 := true
			// 找一个小的
			if nums1[index1] < nums2[index2] {
				m1 = nums1[index1]
				index1++
			} else {
				m1 = nums2[index2]
				index2++
				mIn1 = false
			}
			// 如果是奇数，直接返回
			if total_length%2 != 0 {
				return float64(m1)
			}

			// 如果是偶数，需要再往后找一个数，然后求均值
			m2 := 0
			if mIn1 { // 如果在nums1中找到的m1，那么nums2[index2]一定没超出范围，再取一个较小的
				m2 = nums2[index2]
				if index1 < length1 && nums1[index1] < m2 {
					m2 = nums1[index1]
				}
			} else { // 如果在nums2中找到的m1，那么nums1[index1]一定没超出范围，再取一个较小的
				m2 = nums1[index1]
				if index2 < length2 && nums2[index2] < m2 {
					m2 = nums2[index2]
				}
			}

			return float64(m1+m2) / 2
		}

		// next1,next2是nums1，nums2中的将可能被过滤位置的下标
		// 后移动k/2
		next1 := index1 + k>>1 - 1
		next2 := index2 + k>>1 - 1

		if next1 >= length1 {
			next1 = length1 - 1
		}

		if next2 >= length2 {
			next2 = length2 - 1
		}

		// 过滤掉小的，移动到后一个下标
		// 过滤掉了 next - index + 1 个
		// 如果不到尾部的话， next - index + 1 = (index + k>>1 - 1) - index + 1 = k>>1
		// 所以应该是O(log(m+n))
		// 如果到尾部的话，可以直接算出来了
		if nums1[next1] <= nums2[next2] {
			k -= next1 - index1 + 1
			index1 = next1 + 1
		} else {
			k -= next2 - index2 + 1
			index2 = next2 + 1
		}
	}
}

// 二分法
func findMedianSortedArraysBinary(nums1 []int, nums2 []int) float64 {
	length1, length2 := len(nums1), len(nums2)

	total_length := length1 + length2

	//保证nums1为较短的
	if length1 > length2 {
		nums1, nums2 = nums2, nums1
		length1, length2 = length2, length1
	}
	// 二分法找到刀切合理的位置（找nums1的）
	// 注意，一般二分查找high = length1 - 1，因为过界了
	// 但是此处是可以的，当刀放置到length1出，表示其左边都是小的数字
	low, high := 0, length1

	// nums1和nums2左边的总数目
	left_total := (total_length + 1) / 2

	// mid1, mid2为nums1和nums2中刀切的位置
	mid1, mid2 := 0, 0

	// 在 [low, high] 中切，直到找到合适位置
	for low <= high {
		// | 表示刀的位置，往中间切
		// nums1[low] ... nums1[mid1-1] | nums1[mid1] ... nums1[high - 1]
		//            ... nums2[mid2-1] | nums2[mid2] ...
		mid1 = (low + high) / 2
		mid2 = left_total - mid1

		// mid1 = 0 和 length1都表示nums1已经到边界了，可以算出了
		if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			// nums1刀左边的大于nums2右边的，那么就切多了，在[low, mid1-1]中切
			high = mid1 - 1
		} else if mid1 < length1 && nums1[mid1] < nums2[mid2-1] {
			// nums1刀右边的小于nums2左边的，那么就切少了，在[mid1+1 high]中切
			low = mid1 + 1
		} else {
			break
		}
	}

	m1, m2 := 0, 0
	if mid1 == 0 { // 如果切到nums1最左边，侧表示nums1全部属于大的
		m1 = nums2[mid2-1]
	} else if mid2 == 0 { // 如果切到nums2最左边，侧表示nums2全部属于大的
		m1 = nums1[mid1-1]
	} else { //否则取左边中较大的 (取小的中较大的一个)
		m1 = nums1[mid1-1]
		if m1 < nums2[mid2-1] {
			m1 = nums2[mid2-1]
		}
	}

	if total_length%2 != 0 {
		return float64(m1)
	}

	if mid1 == length1 { // 如果切到nums1最右边，侧表示nums1全部属于小的
		m2 = nums2[mid2]
	} else if mid2 == length2 { // 如果切到nums2最右边，侧表示nums2全部属于小的
		m2 = nums1[mid1]
	} else { //否则取右边中较小的 (取大的中较小的一个)
		m2 = nums1[mid1]
		if m2 > nums2[mid2] {
			m2 = nums2[mid2]
		}
	}

	return float64(m1+m2) / 2
}
