package p49

import "sort"

func groupAnagrams(strs []string) [][]string {
	tempMap := make(map[string][]string)
	var result [][]string
	for i := 0; i < len(strs); i++ {
		key := sortString(strs[i])
		tempMap[key] = append(tempMap[key], strs[i])
	}
	for _, item := range tempMap {
		result = append(result, item)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
