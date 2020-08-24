package leetcode

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	matrixLen := len(matrix)
	if matrixLen == 0 {
		return false
	}
	i := 0
	firstRow := matrix[i]
	j := len(firstRow) - 1
	for i < matrixLen && j >= 0 {
		visit := matrix[i][j]
		if visit > target {
			j--
		} else if visit < target {
			i++
		} else {
			return true
		}
	}
	return false
}
