package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var current []int
	for n := 1; n <= target/candidates[0]; n++ {
		dfs(candidates, target, 0, 0, n, &current, &result)
	}
	return result
}

func dfs(candidates []int, target int, s int, d int, n int, current *[]int, result *[][]int) {
	if d == n {
		if target == 0 {
			*result = append(*result, *current)
			return
		}
	}
	for i := s; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		*current = append(*current, candidates[i])
		dfs(candidates, target-candidates[i], i, d+1, n, current, result)
		*current = (*current)[:len(*current)-1]
	}
}

func main() {
	data := []int{2, 3, 6, 7}
	fmt.Print(combinationSum(data, 7))
}
