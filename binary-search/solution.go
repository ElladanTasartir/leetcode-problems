func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := int((left + right) / 2)
		middleVal := nums[middle]

		if middleVal > target {
			right = middle - 1
			continue
		}

		if middleVal < target {
			left = middle + 1
			continue
		}

		return middle
	}

	return -1
}