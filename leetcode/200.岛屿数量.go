/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	num := 0
	row := len(grid)
	column := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			// 🌟'0'不是0 ！！！
			if grid[i][j] != '0' {
				// DFS(grid, i, j)
				BFS(grid, i, j)
				num++
			}
		}
	}

	return num
}

// 🌟深度搜索：栈 或者 **递归** 都可以，递归写起来简单
func DFS(grid [][]byte, i, j int) {
	row := len(grid)
	column := len(grid[0])

	if i < 0 || i >= row || j < 0 || j >= column {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	DFS(grid, i-1, j)
	DFS(grid, i+1, j)
	DFS(grid, i, j-1)
	DFS(grid, i, j+1)
}

/////////////////////////

// 🌟广度搜索：队列
type pos struct {
	i int
	j int
}

type queue []pos

func (q *queue) push(i, j int) {
	*q = append(*q, pos{
		i: i,
		j: j,
	})
}

func (q *queue) pop() pos {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func BFS(grid [][]byte, i, j int) {
	row := len(grid)
	column := len(grid[0])

	q := new(queue)
	grid[i][j] = '0'
	q.push(i, j)

	for !q.empty() {
		current := q.pop()

		if current.i-1 >= 0 && grid[current.i-1][current.j] != '0' {
			grid[current.i-1][current.j] = '0'
			q.push(current.i-1, current.j)
		}
		if current.i+1 < row && grid[current.i+1][current.j] != '0' {
			grid[current.i+1][current.j] = '0'
			q.push(current.i+1, current.j)
		}
		if current.j-1 >= 0 && grid[current.i][current.j-1] != '0' {
			grid[current.i][current.j-1] = '0'
			q.push(current.i, current.j-1)
		}
		if current.j+1 < column && grid[current.i][current.j+1] != '0' {
			grid[current.i][current.j+1] = '0'
			q.push(current.i, current.j+1)
		}
	}
}

// @lc code=end

