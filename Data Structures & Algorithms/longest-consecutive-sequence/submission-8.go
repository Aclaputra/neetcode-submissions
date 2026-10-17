func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// put into hashmap
	seen := make(map[int]struct{}) 
	for _, num := range nums {
		seen[num] = struct{}{}
	}

	// turn into array
	arr := make([]int, 0)
	for key, _ := range seen {
		arr = append(arr, key)
	}

	// sort it from lowest to the highest
	sort.Ints(arr)

	// traverse it get the total streak
	prev := math.MinInt
	var streak, maxStreak int
	for _, data := range arr {
		if prev == math.MinInt {
			prev = data
			streak = 1
			maxStreak = 1
			continue
		}

		if prev+1 == data {
			prev = data
			streak++
		} else {
			prev = data
			streak = 1
		}

		if streak > maxStreak {
			maxStreak = streak
		}
	}

	return maxStreak
}