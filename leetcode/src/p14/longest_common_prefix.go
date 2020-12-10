package p14

import "strings"

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	// find shortest
	shortestIndex := 0
	for i := 1; i < n; i++ {
		data := strs[i]
		if len(data) < len(strs[shortestIndex]) {
			shortestIndex = i
		}
	}
	// 根据shortest 判断
	result := strings.Builder{}
	target := strs[shortestIndex]
	for i := 0; i < len(target); i++ {
		for _, item := range strs {
			if item[i] != target[i] {
				return result.String()
			}
		}
		result.WriteRune(rune(target[i]))
	}
	return result.String()
}
