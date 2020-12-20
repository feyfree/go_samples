package p30

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	// 字符数组的长度
	n := len(words)
	// 单个字符串的长度
	l := len(words[0])
	expected := make(map[string]int)
	// 初始化 words map
	for i := 0; i < n; i++ {
		expected[words[i]]++
	}
	//
	for i := 0; i <= len(s)-n*l; i++ {
		seen := make(map[string]int)
		count := 0
		for j := 0; j < n; j++ {
			w := s[i+j*l : i+j*l+l]
			it := expected[w]
			if it == 0 {
				break
			}
			seen[w]++
			if seen[w] > it {
				break
			}
			count++
		}
		if count == n {
			result = append(result, i)
		}
	}
	return result
}
