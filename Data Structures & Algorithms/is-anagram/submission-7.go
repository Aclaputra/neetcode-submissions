func isAnagram(s string, t string) bool {
	// it has to be the same length
	if len(s) != len(t) {
		return false
	}

	// turn it into characters array or runes in go
	sRunes, tRunes := []rune(s), []rune(t)
	// sort it
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] > sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] > tRunes[j]
	})

	// if one char is different then it is not an anagram 
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	// its an anagram
	return true
}
