func twoSum(nums []int, target int) []int {
	// Create a map to store the complement of each number
	complementMap := make(map[int]int)

	// Iterate through the array and check if the complement exists in the map
	for i, num := range nums {
		complement := target - num
		if j, ok := complementMap[complement]; ok {
			return []int{j, i}
		}
		complementMap[num] = i
	}

	// No solution found
	return []int{}
}
