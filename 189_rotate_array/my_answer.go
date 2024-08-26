func rotate(nums []int, k int)  {
    // check if k is greater than length of nums
    k = k % len(nums)
    fmt.Println(k)
    // no need to run through array if k is 0
    if k == 0 {
        return
    }

    // create slice that will hold the correct order
    numHolder := make([]int, 0)

    // see where to start rotating
    numPointer := len(nums) - k

    // get the last parts of slice in the front
    numHolder = append(numHolder, nums[numPointer:]...)

    // add the rest of the slice
    numHolder = append(numHolder, nums[:numPointer]...)

    // reassign nums according to correct order help in numHolder
    for idx, _ := range numHolder {
        nums[idx] = numHolder[idx]
    }
}