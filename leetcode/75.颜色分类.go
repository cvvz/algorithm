/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	n := len(nums)
	head, tail := 0, n-1

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[head], nums[i] = nums[i], nums[head]
			head++
		}
	}

	// 注意是从尾向头！
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 2 {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		}
	}
}

// @lc code=end

