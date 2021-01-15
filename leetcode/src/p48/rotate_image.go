package p48

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place,
// which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		temp := matrix[i]
		matrix[i] = matrix[n-1-i]
		matrix[n-1-i] = temp
	}
	// 对角线 交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}
