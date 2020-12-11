package p17

// common solution
func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return make([]string, 0)
	}
	d := []string{" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz"}
	cur := make([]rune, n)
	var ans []string
	dfs(digits, &d, 0, &cur, &ans)
	return ans
}

func dfs(digits string, d *[]string, l int, cur *[]rune, ans *[]string) {
	if l == len(digits) {
		if l > 0 {
			*ans = append(*ans, string(*cur))
		}
		return
	}
	s := (*d)[digits[l]-'0']
	for i := 0; i < len(s); i++ {
		(*cur)[l] = rune(s[i])
		dfs(digits, d, l+1, cur, ans)
	}
}
