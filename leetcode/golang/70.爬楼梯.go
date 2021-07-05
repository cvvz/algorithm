/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	ppre , pre := 0,1
	for i:=1;i<n;i++ {
		ppre , pre = pre, ppre + pre
	}

	return ppre + pre
}

// f(n) = f(n-1) + f(n-2)
// f(1) = 1
// f(2) = 2
// 斐波那契数列
// 时间复杂度：O(n)
// 空间复杂度：O(1)

// @lc code=end

