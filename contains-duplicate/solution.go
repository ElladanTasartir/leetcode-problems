func containsDuplicate(nums []int) bool {
    vMap := map[int]bool{}

    for i := 0; i < len(nums); i++ {
        _, ok := vMap[nums[i]]
        
        if ok {
            return true
        }

        vMap[nums[i]] = true
    }

    return false
}
