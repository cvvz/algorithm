/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var pre,ans []int

	for i:=0;i<=rowIndex;i++ {
		ans = make([]int, i+1)
		ans[0] = 1
		ans[i] = 1
		for j :=1;j<i;j++ {
			ans[j] = pre[j-1] + pre[j]
		}
		pre = ans
	}
	return ans
}
// @lc code=end

