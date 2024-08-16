func majorityElement(nums []int) int {
    // for keeping track of which element has the majority
    // and what the majority number is
    type elementTracker struct {
        majority float64
        elem int
    }

    // hold counts for each element
    elementCount := map[int]int{}

    // put all elements into a map and count them
    for _, element := range nums {
        if _, ok := elementCount[element]; !ok {
            elementCount[element] = 1
        } else {
            elementCount[element] += 1
        }
    }

    tracker := elementTracker{}
    // go through totals of elements and determine which is the majority
    for element, total := range elementCount {
        if float64(total) / 2 > tracker.majority {
            tracker.elem = element
            tracker.majority = float64(total) / 2
        }
    }
    return tracker.elem
}