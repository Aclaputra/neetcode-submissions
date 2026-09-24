func groupAnagrams(strs []string) [][]string {
	// sorting 
	group := make(map[string][]string)
	for _, str := range strs {
		sortedS := sortString(str)
		group[sortedS] = append(group[sortedS], str)
	}

	// traverse the map into result [][]string
	res := make([][]string, 0)
	for _, g := range group {
		res = append(res, g)
	}
	return res
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
