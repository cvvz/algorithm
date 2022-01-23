/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])
	wordLen := len(word)

	visited := make([][]bool, row, row)

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			visited[i] = append(visited[i], false)
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	var dfs func(int, int, int) bool

	dfs = func(i, j, k int) bool {
		if k == wordLen {
			return true
		}

		if i < 0 || j < 0 || i >= row || j >= column {
			return false
		}

		if visited[i][j] {
			return false
		}

		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()

		if board[i][j] != word[k] {
			return false
		}

		for _, direction := range directions {
			if dfs(i+direction[0], j+direction[1], k+1) {
				return true
			}
		}

		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// 🌟🌟🌟
// 1. 用内部函数的方法可以避免参数过多
// 2. 不需要进行字符串的比较，依次比较字符就可以了
// 3. dfs + 回溯

// @lc code=end

