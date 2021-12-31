/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

// 题解中dp的推导过程不太能够理解：
// 为啥每天相对于前一天都进行最优的选择，最终得到的结果就是最优解？
// 为啥局部最优可以等于全局最优？这不就是贪心的思想么。

// @lc code=end

