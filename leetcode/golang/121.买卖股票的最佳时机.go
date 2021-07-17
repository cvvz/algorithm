/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	benefit := 0
	if len(prices) == 0 {
		return benefit
	}

	minPrice := prices[0]

	for i:=1;i<len(prices);i++ {
		if prices[i] - minPrice > benefit {
			benefit = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return benefit
}

// 只用遍历一遍的方法：假设第n天卖，只要找出之前最小值即可
// @lc code=end

