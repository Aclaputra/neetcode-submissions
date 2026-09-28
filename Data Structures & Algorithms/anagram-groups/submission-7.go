func groupAnagrams(strs []string) [][]string {
	// sorting
	group := make(map[string][]string)
	for _, str := range strs {
		sortS := sortString(str)
		group[sortS] = append(group[sortS], str)
	}

	arr := make([][]string, 0, len(strs))
	for _, g := range group {
		arr = append(arr, g)
	}

	return arr
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i,j int) bool {
		return chars[i] > chars[j]
	})
	return string(chars)
}