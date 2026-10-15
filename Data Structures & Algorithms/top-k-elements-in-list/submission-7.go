func topKFrequent(nums []int, k int) []int {
	// create a hashmap to find frequencies easily
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// turn into array then we able to sort it
	arr := make([][2]int, len(freq))
	for key, cnt := range freq {
		arr = append(arr, [2]int{cnt, key})
	}

	// sort the array of frequencies
	sort.Slice(arr, func(i,j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// find the top k
	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = arr[i][1]
	}

	return topK
}
