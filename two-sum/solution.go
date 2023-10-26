func twoSum(nums []int, target int) []int {
    sumsMap := make(map[int]int)

    for i, num := range nums {
        diff := target - num
        if _, ok := sumsMap[diff]; ok {
            return []int{sumsMap[diff], i}
        }
        
        sumsMap[num] = i
    }

    return nil
}
