/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}

	subAns := []int{}
	visited := make([]bool, n, n)

	var backtracking func()
	backtracking = func() {
		if len(subAns) == n {
			// 写法1
			// temp := make([]int, n, n)
			// copy(temp, subAns)
			// ans = append(ans, temp)

			// 写法2
			ans = append(ans, append([]int{}, subAns...))
			return
		}

		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			subAns = append(subAns, nums[i])
			visited[i] = true

			backtracking()

			visited[i] = false
			subAns = subAns[:len(subAns)-1]
		}
	}

	backtracking()

	return ans
}

// @lc code=end

