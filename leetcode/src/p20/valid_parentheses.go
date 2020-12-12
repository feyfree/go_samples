package p20

import "container/list"

func isValid(s string) bool {
	stack := list.New()
	cMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{'}
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if _, ok := cMap[c]; !ok {
			//不存在
			stack.PushBack(c)
		} else if stack.Len() == 0 || cMap[c] != stack.Remove(stack.Back()) {
			return false
		}
	}
	return stack.Len() == 0
}
