func twoSum(nums []int, target int) []int {
    // brute force
	for i := 0; i < len(nums); i++ {
		left := nums[i]
		for j := i+1; j < len(nums); j++ {
			right := nums[j]
			if left + right == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}
