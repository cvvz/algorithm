/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	subAns := []int{}

	visited := make([]bool, n, n)

	var bt func()

	bt = func() {
		if len(subAns) == n {
			temp := make([]int, n, n)
			copy(temp, subAns)
			ans = append(ans, temp)
			return
		}

		duplicate := make(map[int]bool)
		for i := 0; i < n; i++ {
			if visited[i] || duplicate[nums[i]] {
				continue
			}

			duplicate[nums[i]] = true

			visited[i] = true
			subAns = append(subAns, nums[i])

			bt()

			visited[i] = false
			subAns = subAns[:len(subAns)-1]
		}

	}

	bt()

	return ans
}

// @lc code=end

