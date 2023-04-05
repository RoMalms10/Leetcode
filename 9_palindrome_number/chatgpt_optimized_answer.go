func isPalindrome(x int) bool {
	convertedX := strconv.Itoa(x)
	length := len(convertedX)
	mid := length / 2

	j := length - 1
	for i := 0; i < mid; i++ {
		if convertedX[i] != convertedX[j] {
			return false
		}
		j--
	}
	return true
}
