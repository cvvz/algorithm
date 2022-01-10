/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	benefit, lowestPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-lowestPrice > benefit {
			benefit = prices[i] - lowestPrice
		}
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
		}
	}

	return benefit
}

// @lc code=end

