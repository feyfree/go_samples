package p26

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return n
	}
	// 起点是 0 但是实际上表示是len 所以返回index + 1
	index := 0
	for i := 0; i < n; i++ {
		if nums[i] != nums[index] {
			index += 1
			nums[index] = nums[i]
		}
	}
	return index + 1
}
