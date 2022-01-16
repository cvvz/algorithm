/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start

func numIslands(grid [][]byte) int {
	row, column := len(grid), len(grid[0])
	count := 0
	q := &pointQ{}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				count++
				// DFS(grid, i, j)
				// BFS
				func() {
					grid[i][j] = '0'
					q.push(&point{i, j})

					for !q.empty() {
						pos := q.pop()
						r, c := pos.i, pos.j
						if r-1 >= 0 && grid[r-1][c] == '1' {
							grid[r-1][c] = '0'
							q.push(&point{r - 1, c})
						}
						if r+1 < row && grid[r+1][c] == '1' {
							grid[r+1][c] = '0'
							q.push(&point{r + 1, c})
						}
						if c-1 >= 0 && grid[r][c-1] == '1' {
							grid[r][c-1] = '0'
							q.push(&point{r, c - 1})
						}
						if c+1 < column && grid[r][c+1] == '1' {
							grid[r][c+1] = '0'
							q.push(&point{r, c + 1})
						}
					}
				}()
			}
		}
	}

	return count
}

type point struct {
	i, j int
}

type pointQ []*point

func (q *pointQ) push(p *point) {
	*q = append(*q, p)
}

func (q *pointQ) pop() *point {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *pointQ) empty() bool {
	return len(*q) == 0
}

func DFS(grid [][]byte, i, j int) {
	grid[i][j] = '0'

	row, column := len(grid), len(grid[0])
	if i+1 < row && grid[i+1][j] == '1' {
		DFS(grid, i+1, j)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		DFS(grid, i-1, j)
	}
	if j+1 < column && grid[i][j+1] == '1' {
		DFS(grid, i, j+1)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		DFS(grid, i, j-1)
	}
}

// @lc code=end

