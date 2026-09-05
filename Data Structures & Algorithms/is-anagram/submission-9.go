func isAnagram(s string, t string) bool {
	// check length
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	// count characters 
	for i := 0; i < len(s); i++ {
		countS[rune(s[i])]++
		countT[rune(t[i])]++
	}

	// check if its different
	for key, val := range countS {
		if countT[key] != val { // the count of the character is different def not an anagra
			return false
		}
	}
	return true
}
