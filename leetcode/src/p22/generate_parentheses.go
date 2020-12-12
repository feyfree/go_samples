package p22

// 实际上是dfs搜索的问题
func generateParenthesis(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}
	left := n
	right := n
	var cur []rune
	var ans []string
	dfs(&ans, left, right, n, &cur)
	return ans
}

func dfs(ans *[]string, left int, right int, n int, cur *[]rune) {
	if left == 0 && right == 0 {
		*ans = append(*ans, string(*cur))
		return
	}
	if left != 0 {
		cur1 := append(*cur, '(')
		dfs(ans, left-1, right, n, &cur1)
	}
	if right > left {
		cur2 := append(*cur, ')')
		dfs(ans, left, right-1, n, &cur2)
	}
	return
}
