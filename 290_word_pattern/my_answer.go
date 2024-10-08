func wordPattern(pattern string, s string) bool {
	// can't make a pattern if the length is 0
	if len(s) == 0 {
		return false
	}

	// Create maps for both directions so you don't double assign one of them
	charMap := map[byte]string{}
	wordMap := map[string]byte{}

	// split the string into words
	split := strings.Split(s, " ")

	// if the length of the pattern doesn't equal the number of words
	// then it can't make a pattern, so return false
	if len(pattern) != len(split) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		word := split[i]
		char := pattern[i]

		// Check word mapping {dog = a}
		if val, ok := wordMap[word]; ok {
			if val != char {
				return false
			}
		} else {
			wordMap[word] = char
		}

		// Check char mapping {a = dog}
		if val, ok := charMap[char]; ok {
			if val != word {
				return false
			}
		} else {
			charMap[char] = word
		}
	}

	return true
}