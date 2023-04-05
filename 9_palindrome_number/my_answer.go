func isPalindrome(x int) bool {
	convertedX := strconv.Itoa(x)

	j := len(convertedX) - 1
	for i := 0; i < j; i++ {
		if convertedX[i] != convertedX[j] {
			return false
		}
		j--
	}
	return true
}