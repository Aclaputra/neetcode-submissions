func hasDuplicate(nums []int) bool {
	// if value more than once in the array then true
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	} 
	return len(seen) < len(nums)
}
