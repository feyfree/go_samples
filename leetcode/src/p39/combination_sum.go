package p39

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var current []int
	for n := 1; n <= target/candidates[0]; n++ {
		dfs(candidates, target, 0, 0, n, current, &result)
	}
	return result
}

// slice 可以认为是一个结构体, 里面包含了一个数组的指针
func dfs(candidates []int, target int, s int, d int, n int, current []int, result *[][]int) {
	if d == n {
		if target == 0 {
			*result = append(*result, append([]int(nil), current...))
			return
		}
	}
	for i := s; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		// 这个地方有点坑 需要注意一下
		dfs(candidates, target-candidates[i], i, d+1, n, append(current, candidates[i]), result)
	}
}
