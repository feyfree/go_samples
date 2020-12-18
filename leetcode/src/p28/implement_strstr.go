package main

import (
	"fmt"
	"math"
)

func main() {
	hayStack := "hello"
	needle := "ll"
	fmt.Println(strStr(hayStack, needle))
}

func strStr(haystack string, needle string) int {
	// 使用boyer-moore 算法实现
	return boyerMoore(haystack, needle)
}

func boyerMoore(hayStack string, needle string) int {
	m := len(hayStack)
	n := len(needle)
	switch {
	case n == 0:
		return 0
	case n > m:
		return -1
	default:
		// 构造右边数组  实际上 index 表示是rune 的相对位置
		// 右边数组表示的是 pattern 模块最右边出现的位置
		right := [26]int{}
		for i := 0; i < 26; i++ {
			right[i] = -1
		}
		for i := 0; i < n; i++ {
			right[rune(needle[i])-'a'] = i
		}
		var skip int
		for i := 0; i <= m-n; i += skip {
			skip = 0
			for j := n - 1; j >= 0; j-- {
				if needle[j] != hayStack[i+j] {
					skip = int(math.Max(1, float64(j-right[hayStack[i+j]-'a'])))
					break
				}
			}
			if skip == 0 {
				return i
			}
		}
		return -1
	}
}

// kmp 算法实现
func kmp(hayStack string, needle string) int {
	return 0
}
