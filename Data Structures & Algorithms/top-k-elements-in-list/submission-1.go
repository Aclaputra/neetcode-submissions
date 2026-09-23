func topKFrequent(nums []int, k int) []int {
	// count frequency to map
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// put in array so it could be sorted
	arr := make([][2]int, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	// sort
	sort.Slice(arr, func(i,j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// top k frequent
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}
