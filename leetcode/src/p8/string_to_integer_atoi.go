package p8

import (
	"math"
)

//
func myAtoi(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	// 有效数字 在 s中的开始索引
	cursor := 0
	// 返回结果
	result := 0
	// 符号位
	sign := 1
	// 跳过空字符串
	for cursor < n && s[cursor] == ' ' {
		cursor++
	}
	// 获取符号位
	if cursor < n && (s[cursor] == '+' || s[cursor] == '-') {
		if s[cursor] == '+' {
			sign = 1
		}
		if s[cursor] == '-' {
			sign = -1
		}
		cursor++
	}
	for cursor < n {
		digit := s[cursor] - '0'
		if digit < 0 || digit > 9 {
			break
		}
		// 判断溢出条件
		overflow := math.MaxInt32/10 < result || (math.MaxInt32/10 == result && math.MaxInt32%10 < digit)
		if overflow {
			if sign > 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + int(digit)
		cursor++
	}
	return result * sign
}
