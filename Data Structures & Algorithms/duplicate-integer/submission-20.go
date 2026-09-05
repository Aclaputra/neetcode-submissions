func hasDuplicate(nums []int) bool {
	// hashmap 
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	
	return false
}
