package p3

import (
	"math"
	"strings"
)

// 线性扫描算法 keep track of every character
func solution(s string) int {
	n := len(s)
	ans := 0
	// 生成s的切片
	sli := s[0:0]
	// i 相当于 start, j 相当于end , end++ 作为循环
	for i, j := 0, 0; j < n; j++ {
		// 查找切片中 s[j] 的位置
		if index := strings.IndexByte(sli, s[j]); index != -1 {
			i += index + 1
		}
		// 切片更新
		sli = s[i : j+1]
		// 计算结果
		ans = int(math.Max(float64(ans), float64(j-i+1)))
	}
	return ans
}
