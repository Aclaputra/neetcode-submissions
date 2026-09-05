func isAnagram(s string, t string) bool {
	// anagrams have to be the same length
	if len(s) != len(t) {
		return false
	}

	// sort characters
	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// have the same length, iterate and compare
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}
