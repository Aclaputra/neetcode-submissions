func topKFrequent(nums []int, k int) []int {
	// sorting method
	// count frequent by map
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// put into array to sort
	arr := make([][2]int, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	// sort the data
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	// traverse the data amount of k append to result
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}
