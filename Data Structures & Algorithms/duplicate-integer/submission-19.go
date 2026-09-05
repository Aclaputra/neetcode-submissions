func hasDuplicate(nums []int) bool {
	// sort 
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		left := nums[i-1]
		right := nums[i]
		if left == right {
			return true
		}
	}
	return false
}
