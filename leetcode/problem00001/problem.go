package problem00001

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int)

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
