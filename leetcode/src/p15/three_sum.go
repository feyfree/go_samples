package p15

import "sort"

//Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
//Find all unique triplets in the array which gives the sum of zero.
//
//Notice that the solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return make([][]int, 0)
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := n - 1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if start-1 > i && nums[start] == nums[start-1] {
				start++
				continue
			}
			if sum == 0 {
				temp := make([]int, 3)
				temp[0] = nums[i]
				temp[1] = nums[start]
				temp[2] = nums[end]
				result = append(result, temp)
				start++
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
