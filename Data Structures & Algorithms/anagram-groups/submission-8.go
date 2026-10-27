func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	// sort the string put in the hashmap for grouping by key, value is original str
	for _, str := range strs {
		anagram := sortString(str)
		groups[anagram] = append(groups[anagram], str)
	}

	res := make([][]string, 0)
	// turn into 2d string array
	for _, value := range groups {
		res = append(res, value)
	}

	return res
}

func sortString(str string) string {
	arr := []rune(str)
	sort.Slice(arr, func(i,j int) bool {
		return arr[i] > arr[j]
	})

	return string(arr)
}
