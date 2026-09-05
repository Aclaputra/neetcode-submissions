func hasDuplicate(nums []int) bool {
	// if value more than once in the array then true
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			// duplicate
			return true
		}
		seen[num] = true
	}
	return false
}
