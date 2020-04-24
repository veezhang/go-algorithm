package problem00011

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 双指针
// 时间复杂度： O(n)
// 空间复杂度： O(1)
func maxArea(height []int) int {
	area := 0

	for left, right := 0, len(height)-1; left < right; {
		area = max(area, (right-left)*min(height[left], height[right]))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}
