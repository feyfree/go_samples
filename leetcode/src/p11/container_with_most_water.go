package p11

import "math"

func maxArea(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	start := 0
	end := n - 1
	capacity := 0.0
	for start <= end {
		capacity = math.Max((float64(end)-float64(start))*math.Min(float64(height[start]), float64(height[end])), capacity)
		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}
	return int(capacity)
}
