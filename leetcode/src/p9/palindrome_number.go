package p9

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	// 数字实际上就是栈的思想 反转 实际上就是出栈
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return x == reverse || x == reverse/10
}
