func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// sort
	sRune := []rune(s)
	tRune := []rune(t)
	sort.Slice(sRune, func(i,j int) bool {
		return sRune[i] > sRune[j]
	})
	sort.Slice(tRune, func(i,j int) bool {
		return tRune[i] > tRune[j]
	})

	// traverse the T check if its an anagram of S
	for i := 0; i < len(s); i++ {
		if tRune[i] != sRune[i] {
			return false
		}
	}

	return true
}
