package two_sum

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mIndex, ok := m[target-nums[i]]
		if ok {
			return []int{mIndex, i}
		}
		m[nums[i]] = i
	}

	return []int{-1, -1}
}