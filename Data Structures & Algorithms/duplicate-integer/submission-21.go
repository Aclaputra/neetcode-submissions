func hasDuplicate(nums []int) bool {
	// hash map with struct
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	}

	return len(seen) < len(nums)
}
