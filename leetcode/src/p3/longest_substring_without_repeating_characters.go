package p3

import "math"

// 线性扫描算法 keep track of every character
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	// 返回的结果
	result := 0
	// last index of each character
	// here we user 256 as character set total num
	lastIndexArray := [256]int{}
	for i := range lastIndexArray {
		lastIndexArray[i] = -1
	}
	start := 0
	// 设置滑动窗口
	for end := 0; end < n; end++ {
		start = int(math.Max(float64(start), float64(lastIndexArray[s[end]]+1)))
		result = int(math.Max(float64(result), float64(end-start+1)))
		lastIndexArray[s[end]] = end
	}
	return result
}
