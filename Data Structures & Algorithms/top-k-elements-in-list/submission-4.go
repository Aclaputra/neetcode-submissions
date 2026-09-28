func topKFrequent(nums []int, k int) []int {
	// count frequencies with map
	frequent := make(map[int]int)
	for _, num := range nums {
		frequent[num]++
	}

	// turn it into array to be able to sort
	arr := make([][2]int, 0, len(frequent))
	for num, cnt := range frequent {
		arr = append(arr, [2]int{cnt, num})
	}

	// sort the array desc
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})	

	// traverse the array to find top k frequent
	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = arr[i][1]
	}

	return topK
}
