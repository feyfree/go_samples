package p27

func removeElement(nums []int, val int) int {
	index := 0
	// 这里的index 跟 p26 index 不同, index 最后表示的是next place index 实际上就是len
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
