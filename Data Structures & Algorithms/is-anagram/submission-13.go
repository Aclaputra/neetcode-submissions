func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// sorting
	charS := []rune(s)
	charT := []rune(t)
	sort.Slice(charS, func(i,j int) bool {
		return charS[i] < charS[j]
	})	
	sort.Slice(charT, func(i,j int) bool {
		return charT[i] < charT[j]
	})	

	return string(charS) == string(charT)
}
