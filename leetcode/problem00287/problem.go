package problem00287

func findDuplicate1(nums []int) int {
	target := -1
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // 防止越界

		cnt := 0
		// 获取 nums 数组中小于等于 mid 的个数
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		if cnt <= mid { // 左边没有多余的数字，就表示不存在，得在右边找
			left = mid + 1
		} else { // 右边有多余的数字，记录下这个数，往左边找找看，可能就是这个数，也可能还在左边
			right = mid - 1
			target = mid
		}
	}

	return target
}

func findDuplicate2(nums []int) int {
	target := 0
	length := len(nums) // length 为 n + 1 ，n := length - 1

	bitMax := 31
	for 0 == (length-1)>>bitMax {
		bitMax--
	}

	for bit := 0; bit <= bitMax; bit++ {
		x, y := 0, 0
		for i := 0; i < length; i++ {
			if nums[i]&(1<<bit) > 0 {
				x++
			}
			// [1,n] -> [1,length)
			// 共用下循环
			// i = 0 的时候下面的条件也不成立 if i >= 1 && i&(1<<bit) > 0 {} ==> if i&(1<<bit) > 0 {}
			if i >= 1 && i&(1<<bit) > 0 {
				y++
			}
		}

		if x > y {
			target |= 1 << bit
		}
	}

	return target
}

func findDuplicate3(nums []int) int {
	slow, fast := 0, 0

	// 假设环之前需要 m 步， 整个环需要 c 步
	// 假设相遇的时候， slow 移动了 n 步， 则 fast 移动 2n 步，比 slow 多移动 n 步，则 n % c == 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	// 相遇的时候 slow 在环中移动了 n - m 步， 再让 slow 移动 m 步骤，则 show 在环中移动了 n 步
	// 是环的整数被，又回到到了环的起始点
	// 然后让 finder 陪着一起走，则相遇的时候正好在环的起始点
	finder := 0
	for finder != slow {
		finder = nums[finder]
		slow = nums[slow]
	}
	return finder
}
