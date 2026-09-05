func hasDuplicate(nums []int) bool {
	// if value more than once in the array then true
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		left := nums[i-1]
		right := nums[i]
		if left == right {
			// duplicate
			return true
		}
	}
	return false
}
