package p5

func elegantSolution(s string) string {
	if len(s) == 0 {
		return ""
	}
	// dp[i][j] means does string from position i to j is palindrome
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	left, length := 0, 1
	for j := 0; j < len(s); j++ {
		dp[j][j] = true
		for i := 0; i < j; i++ {
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > length {
				length = j - i + 1
				left = i
			}
		}
	}
	return s[left : left+length]
}
