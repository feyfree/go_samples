package p68

import "strings"

// round-robin 算法
func solution(words []string, maxWidth int) []string {
	var result []string
	if len(words) == 0 {
		return result
	}
	// 当前行
	var current []string

	// 当前行拼接起来后的长度
	curSize := 0

	// 索引
	index := 0
	for index < len(words) {
		wordSize := len(words[index])
		if len(current) == 0 || curSize+wordSize < maxWidth {
			current = append(current, words[index])
			if len(current) == 1 {
				curSize += wordSize
			} else {
				curSize += wordSize + 1
			}
			index++
		} else {
			idx := 0
			for curSize < maxWidth {
				current[idx] = current[idx] + " "
				curSize++
				idx++
				if idx >= len(current)-1 {
					idx = 0
				}
			}
			result = append(result, strings.Join(current, " "))
			current = nil
			curSize = 0
		}
	}

	if len(current) > 0 {
		last := strings.Builder{}
		last.WriteString(strings.Join(current, " "))
		for curSize < maxWidth {
			last.WriteByte(' ')
			curSize++
		}
		result = append(result, last.String())
	}

	return result
}
