func isAnagram(s string, t string) bool {
	// check the same length
	if len(s) != len(t) {
		return false
	}
	// sort both words to runes
	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})
	// check the runes array is the same after sorted
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}
