func isAnagram(s string, t string) bool {
	// check the length
	if len(s) != len(t) {
		return false
	}

	// hashmap count
	countS, countT := make(map[rune]int), make(map[rune]int)
	// map the both inputs and count it
	for i := 0; i < len(s); i++ {
		countS[rune(s[i])]++
		countT[rune(t[i])]++
	}

	// check if there is different in count
	for key, val := range countS {
		if countT[key] != val {
			return false
		}
	}
	return true
}
