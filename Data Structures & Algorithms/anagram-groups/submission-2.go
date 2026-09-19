func groupAnagrams(strs []string) [][]string {
	// create hashmap grouping a sorted string
	res := make(map[string][]string)

	for idx, s := range strs {
		sortedS := sortString(s)
		res[sortedS] = append(res[sortedS], strs[idx])
	}

	// map the hashmap to 2d array as result
	var result [][]string
	for _, value := range res {
		result = append(result, value)
	}

	return result
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
