package p32

import (
	"container/list"
	"math"
)

func longestValidParentheses(s string) int {
	result := 0
	if len(s) == 0 {
		return result
	}
	start := 0
	stack := list.New()
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c == '(' {
			stack.PushBack(i)
		} else {
			if stack.Len() == 0 {
				start = i + 1
			} else {
				element := stack.Back()
				stack.Remove(element)
				var temp int
				if stack.Len() == 0 {
					temp = i - start + 1
				} else {
					temp = i - stack.Back().Value.(int)
				}
				result = int(math.Max(float64(temp), float64(result)))
			}
		}
	}
	return result
}
