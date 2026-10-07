func topKFrequent(nums []int, k int) []int {
	// objective: find frequencies then find the top k
	// traverse array to find frequency using hashmap
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// turn into array to sort
	arr := make([][2]int, len(freq))
	for key, cnt := range freq {
		arr = append(arr, [2]int{cnt, key})
	}

	// sort
	sort.Slice(arr, func(i,j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// traverse find top k
	topK := make([]int, k)
	for i:=0; i < k; i++ {
		topK[i] = arr[i][1]
	}

	return topK
}
