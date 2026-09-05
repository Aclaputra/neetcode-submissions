func hasDuplicate(nums []int) bool {
	duplicateFinder := make(map[int]int, 0)
	for _, num := range nums {
		duplicateFinder[num]++
		if duplicateFinder[num] > 1 {
			return true
		}
	}   

	return false
}
