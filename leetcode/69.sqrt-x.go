/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		mid2 := mid * mid
		if mid2 <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// @lc code=end

