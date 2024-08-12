func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums3 := make([]int, m)
    copy(nums3, nums1[0:m])
    i := 0
    j := 0
    finalPointer := 0
    for i < m {
        // if already looped through nums2 all the way
        if j >= n {
            nums1[finalPointer] = nums3[i]
            i++
            finalPointer++
            continue            
        }
        // make switch for which value is smaller
        switch nums3[i] > nums2[j] {
            case true:
                nums1[finalPointer] = nums2[j]
                j++
            case false:
                nums1[finalPointer] = nums3[i]
                i++
        }
        finalPointer++
    }
    // if there are any values left in nums2
    for j < n {
        nums1[finalPointer] = nums2[j]
        j++
        finalPointer++
    }
}