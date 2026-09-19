func groupAnagrams(strs []string) [][]string {
	// hashmap grouping sorted string
	res := make(map[string][]string)

	for id, s := range strs {
		sortedS := sortString(s)
		res[sortedS] = append(res[sortedS], strs[id])
	}

	var result [][]string
	for _, value := range res {
		result = append(result, value)
	}

	return result
}

func sortString(str string) string {
	characters := []rune(str)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
