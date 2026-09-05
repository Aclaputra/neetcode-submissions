func isAnagram(s string, t string) bool {
	// sorting
	// it has to be the same length
	if len(s) != len(t) {
		return false
	}

	// sorting
	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// sorted
	// check if its different then its not an anagram
	for i := 0; i < len(sRunes); i++ { // len(sRunes) cause there is already validatio of length up there
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}
