func groupAnagrams(strs []string) [][]string {
	// create hashmap to groping sorted string
	res := make(map[string][]string)

	for id, s := range strs {
		sortedS := sortString(s)
		res[sortedS] = append(res[sortedS], strs[id])
	}

	// convert the hashmap into 2d array for result
	var result [][]string
	for _, val := range res {
		result = append(result, val)
	}

	return result
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i,j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

