package p47

import "sort"

// Given a collection of numbers, nums,
// that might contain duplicates,
// return all possible unique permutations in any order.

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var result [][]int
	var current []int
	dfs(nums, current, used, &result)
	return result
}

func dfs(nums []int, current []int, used []bool, result *[][]int) {
	if len(current) == len(nums) {
		*result = append(*result, append([]int(nil), current...))
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		dfs(nums, append(current, nums[i]), used, result)
		used[i] = false
	}
}
