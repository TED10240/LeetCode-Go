package leetcode

func rotate(matrix [][]int) {
	length := len(matrix)
	//rotate by diagonal 对角线变化
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//rotate by vertical centerLine 竖直轴对称翻转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
