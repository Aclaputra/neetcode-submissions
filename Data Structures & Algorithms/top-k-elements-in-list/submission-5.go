func topKFrequent(nums []int, k int) []int {
	// count frequencies with map
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// turn it into array to sort
	arr := make([][2]int, 0, len(freq))
	for val, count  := range freq {
		arr = append(arr, [2]int{count, val})
	}

	// sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// traverse the sortedArr to find Top K
	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = arr[i][1]
	}

	return topK
}
