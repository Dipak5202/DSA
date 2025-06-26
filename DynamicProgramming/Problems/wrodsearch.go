package dynamicprogramming

// word search
// Given a 2D board and a word, find if the word exists in the grid
func WordSearch(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, word, 0, i, j, visited) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, word string, index, row, col int, visited [][]bool) bool {
	if index == len(word) {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || visited[row][col] || board[row][col] != word[index] {
		return false
	}

	visited[row][col] = true

	found := dfs(board, word, index+1, row+1, col, visited) ||
		dfs(board, word, index+1, row-1, col, visited) ||
		dfs(board, word, index+1, row, col+1, visited) ||
		dfs(board, word, index+1, row, col-1, visited)

	visited[row][col] = false // backtrack
	return found
}
