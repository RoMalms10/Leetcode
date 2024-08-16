func removeDuplicates(nums []int) int {
    numMap := map[int]int{}
    numOrder := make([]int, 0)
    for _, num := range nums {
        if _, ok := numMap[num]; ok && numMap[num] == 1 {
            numMap[num] = 2
            numOrder = append(numOrder, num)
        }
        if _, ok := numMap[num]; !ok {
            numMap[num] = 1
            numOrder = append(numOrder, num)
        }
    }
    for i, num := range numOrder {
        nums[i] = num
    }
    return len(numOrder)
}