package p1

//
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		var anotherIndex = -1
		for j := len(nums) - 1; j > i; j-- {
			if another == nums[j] {
				anotherIndex = j
			}
		}
		if anotherIndex > 0 {
			result := []int{i, anotherIndex}
			return result
		}
	}
	return nil
}
