func removeDuplicates(nums []int) int {
	numberCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		number := nums[i]
		if _, ok := numberCount[number]; !ok {
			numberCount[number] = 1
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	return len(numberCount)
 }