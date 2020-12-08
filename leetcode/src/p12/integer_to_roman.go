package p12

import "strings"

// 罗马数字
func intToRoman(num int) string {
	values := []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC",
		"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM",
		"M", "MM", "MMM"}
	numArray := [4]int{}
	i := 3
	for num > 0 {
		numArray[i] = num % 10
		num = num / 10
		i--
	}
	result := strings.Builder{}
	for j := 0; j < 4; j++ {
		if numArray[j] != 0 {
			result.WriteString(values[9*(3-j)+numArray[j]-1])
		}
	}
	return result.String()
}
