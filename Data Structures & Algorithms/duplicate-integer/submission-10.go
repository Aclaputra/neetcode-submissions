func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = struct{}{}
	}
	return len(seen) < len(nums)
}
