func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	findAnagram := make(map[rune]int)
	for _, data := range s {
		findAnagram[data]++
	}

	for _, data := range t {
		findAnagram[data]--
		if findAnagram[data] < 0 {
			return false
		}
	}

	return true
}
