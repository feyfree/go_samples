package p3

import "testing"

func TestSolution(t *testing.T) {
	result := lengthOfLongestSubstring(" ")
	t.Log(result)
}

func TestSlice(t *testing.T) {
	//array := [...]int{1, 2, 3}
	//slice := array[0:0]
	//t.Logf("%p", &array)
	//t.Log(cap(slice))
	//t.Logf("%p", &slice[0])

	//

	a := "123"
	part := a[0:0]
	t.Log(part)
	t.Logf("%p", &a)
	t.Logf("%p", &part)
}
