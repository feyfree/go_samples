package p5

//  最长回文子字符串 动态规划
func longestPalindrome(s string) string {
	var n = len(s)
	if n <= 0 {
		return ""
	}
	// 动态规划 -- 初始化dp  dp[i][j] = true 表示
	//start at i and end at j 是有效的
	var dp [][]bool
	for i := 0; i < n; i++ {
		dp = append(dp, []bool{})
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], i == j)
		}
	}
	var maxLen = 1
	var begin = 0
	// 状态定义
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[j] != s[i] {
				dp[i][j] = false
			} else {
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && (j-i+1) > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
