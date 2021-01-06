package p46

// Given an array nums of distinct integers,
// return all the possible permutations.
// You can return the answer in any order.
func permute(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	var current []int
	dfs(nums, 0, current, &result, visited)
	return result
}

func dfs(nums []int, d int, current []int, result *[][]int, visited []bool) {
	if d == len(nums) {
		*result = append(*result, append([]int{nil}, current...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		dfs(nums, d+1, append(current, nums[i]), result, visited)
		visited[i] = false
	}
}
