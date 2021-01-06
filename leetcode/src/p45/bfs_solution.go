package p45

import (
	"container/list"
)

// Runtime: 256 ms
// Memory Usage: 6.5 MB
func solution(nums []int) int {
	n := len(nums)
	marked := make([]bool, n)
	dist := make([]int, n)
	bfs(nums, marked, dist)
	return dist[n-1]
}

func bfs(nums []int, marked []bool, dist []int) {
	queue := list.New()
	marked[0] = true
	queue.PushBack(0)
	for queue.Len() > 0 {
		e := queue.Remove(queue.Front())
		s := e.(int)
		step := nums[s]
		for i := s + 1; i <= s+step && i < len(nums); i++ {
			if !marked[i] {
				dist[i] = dist[s] + 1
				marked[i] = true
				queue.PushBack(i)
			}
			if i == len(nums)-1 {
				break
			}
		}
	}
}
