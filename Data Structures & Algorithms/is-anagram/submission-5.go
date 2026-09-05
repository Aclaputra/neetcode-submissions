func isAnagram(s string, t string) bool {
	// check the same length 
	if len(s) != len(t) {
		return false
	}
	// sort both of the strings to rune array
	sRunes, tRunes := []rune(s), []rune(t)
	bothLength := len(sRunes)
	sort.Slice(sRunes, func(i, j int) bool{
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})
	// check if both of the sorted rune array has the same value 
	for i := 0; i < bothLength; i++ {
		if sRunes[i] != tRunes[i] {
			// not anagram
			return false
		}
	}

	// is anagram
	return true
}
