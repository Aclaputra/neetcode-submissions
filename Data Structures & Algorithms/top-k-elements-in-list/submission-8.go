func topKFrequent(nums []int, k int) []int {
	// find frequencies with hashmap (since its in any order)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// transform into array to sort frequencies
	arr := make([][2]int, 0)
	for num, count := range freq {
		arr = append(arr, [2]int{count, num})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// traverse the top K then return
	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = arr[i][1]
	}

	return topK
}
