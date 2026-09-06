func isAnagram(s string, t string) bool {
	// hash table
	// check length
	if len(s) != len(t) {
		return false
	}

	// check count by hash table
	// s and t consist of lowercase English letters. so we use lowercase 
	// alphabets has 26 characters
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i] - 'a']++
		count[t[i] - 'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}
