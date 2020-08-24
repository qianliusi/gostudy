package leetcode

import "fmt"

func Exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, word) {
				return true
			}
		}
	}
	return false
}
func dfs(b [][]byte, i int, j int, w string) bool {
	if (i < 0 || i >= len(b) || j < 0 || j >= len(b[i])) && len(w) > 0 {
		return false
	}
	if len(w) == 0 {
		return true
	}
	fmt.Println(string(b[i][j]))
	if b[i][j] != w[0] {
		return false
	}
	t := b[i][j]
	b[i][j] = '0'
	s := w[1:]
	if dfs(b, i-1, j, s) || dfs(b, i+1, j, s) || dfs(b, i, j-1, s) || dfs(b, i, j+1, s) {
		b[i][j] = t
		return true
	}
	b[i][j] = t
	return false
}
