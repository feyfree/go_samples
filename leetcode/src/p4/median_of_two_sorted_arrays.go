package p4

import "math"

// 寻找两个数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	// 保证两个数组的都是length小的作为first argument
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	mid := (n1 + n2 + 1) / 2
	l := 0
	r := n1
	for l < r {
		// 二分法定位 m1 的位置
		m1 := l + (r-l)/2
		// m2 的位置
		m2 := mid - m1
		// 第一个数组取少了
		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			// 第一个数组取多了
			r = m1
		}
	}
	// 最终的索引
	m1 := l
	m2 := mid - l
	// num1[m1 -1] 和 num2[m2 -1] 的最大值
	var min1 int
	var min2 int
	// 边界条件
	if m1 <= 0 {
		min1 = math.MinInt32
	} else {
		min1 = nums1[m1-1]
	}

	if m2 <= 0 {
		min2 = math.MinInt32
	} else {
		min2 = nums2[m2-1]
	}
	c1 := math.Max(float64(min1), float64(min2))
	// 判断奇偶性的 位运算， 相当于 % 2
	if (n2+n1)&1 == 1 {
		return c1
	}
	// 如果是偶数
	// num1[m1] 和 num2[m2] 的最小值
	var max1 int
	var max2 int
	if m1 >= n1 {
		max1 = math.MaxInt32
	} else {
		max1 = nums1[m1]
	}

	if m2 >= n2 {
		max2 = math.MaxInt32
	} else {
		max2 = nums2[m2]
	}
	c2 := math.Min(float64(max1), float64(max2))
	return (c1 + c2) * 0.5

}
