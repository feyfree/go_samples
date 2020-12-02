package p4

import "testing"

func TestSolution(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	t.Log(result)
}
