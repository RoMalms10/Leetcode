func removeElement(nums []int, val int) int {
	pointerOne := 0
	pointerTwo := 1
	totalNotNum := 0
	for pointerTwo < len(nums) {
		fmt.Println("start of loop", nums)
		if nums[pointerOne] == val && nums[pointerTwo] != val {
			nums[pointerOne] = nums[pointerTwo]
			nums[pointerTwo] = val
		}
		if nums[pointerOne] == val && nums[pointerTwo] == val {
			for pointerThree := pointerTwo + 1; pointerThree < len(nums); pointerThree++ {
				if nums[pointerThree] != val {
					nums[pointerOne] = nums[pointerThree]
					nums[pointerThree] = val
					break
				}
			}
		}
		pointerOne++
		pointerTwo++
		fmt.Println("end of loop", nums)
	}
	for _, num := range nums {
		if num != val {
			totalNotNum++
		}
	}
	return totalNotNum
}