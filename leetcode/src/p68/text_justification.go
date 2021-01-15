package p68

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	if len(words) == 0 {
		return result
	}
	left := 0
	right := 0
	for left < len(words) {
		size := 0
		for right < len(words) && size+len(words[right])+1 <= maxWidth+1 {
			size += len(words[right]) + 1
			right++
		}
		space := 1
		extra := 0
		if right != len(words) && right != left+1 {
			space = (maxWidth-size+1)/(right-left-1) + 1
			extra = (maxWidth - size + 1) % (right - left - 1)
		}
		sb := strings.Builder{}
		sb.WriteString(words[left])
		left++
		for left < right {
			for i := 0; i < space; i++ {
				sb.WriteByte(' ')
			}
			if extra > 0 {
				sb.WriteByte(' ')
				extra--
			}
			sb.WriteString(words[left])
			left++
		}
		remaining := maxWidth - sb.Len()
		for remaining > 0 {
			sb.WriteByte(' ')
			remaining--
		}
		result = append(result, sb.String())
	}
	return result
}
