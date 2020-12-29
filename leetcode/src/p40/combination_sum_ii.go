package p40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// 首先进行排序
	sort.Ints(candidates)
	var current []int
	var result [][]int
	dfs(candidates, target, 0, current, &result)
	return result
}

//
func dfs(candidates []int, target int, s int, current []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int(nil), current...))
	}
	for i := s; i < len(candidates); i++ {
		c := candidates[i]
		if c > target {
			break
		}
		if i > s && c == candidates[i-1] {
			continue
		}
		dfs(candidates, target-c, i+1, append(current, c), result)
	}
}
