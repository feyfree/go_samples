package p45

import "math"

// 贪心解法
func jump(nums []int) int {
	// 需要跳跃的步数
	steps := 0
	// near 表示上一个点能到达的最远距离
	near := 0
	// 表示 通过 1 个step 可以到达的最远距离
	far := 0
	for i := 0; i < len(nums); i++ {
		if i > near {
			steps++
			// 表示之前能到达的距离
			near = far
		}
		// 在该点的next step 最远的距离
		far = int(math.Max(float64(far), float64(i+nums[i])))
	}
	return steps
}
