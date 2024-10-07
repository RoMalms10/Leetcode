func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) == 0 {
		return false
	}

	charCount := map[rune]int{}
	for _, char := range magazine {
		// loop through magazine and take count of chars
		if _, ok := charCount[char]; !ok {
			charCount[char] = 1
		} else {
			charCount[char] += 1
		}
	}

	// loop through ransomNote and then remove chars from count
	for _, char := range ransomNote {
		// check if it exists
		if _, ok := charCount[char]; !ok {
			return false
		}
		// if the count is already at 0, then that means there are no letters left to use
		if charCount[char] == 0 {
			return false
		}

		charCount[char] -= 1
	}

	return true
}