func hasDuplicate(nums []int) bool {
	// brute force
	for i := 0; i < len(nums); i++ {
		left := nums[i]
		for j := i+1; j < len(nums); j++ {
			right := nums[j]
			
			isDuplicate := left == right
			if isDuplicate {
				return true
			}
		}
	}

	return false
}
