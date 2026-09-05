func hasDuplicate(nums []int) bool {
	// if more than once then true
	for i := 0; i < len(nums); i++ {
		left := nums[i]
		for j := i + 1; j < len(nums); j++ {
			right := nums[j]
			if left == right {
				return true
			}
		}
	}
	return false
}
