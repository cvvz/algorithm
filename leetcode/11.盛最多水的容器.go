/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxArea := 0
	for start != end {
		area := min(height[start], height[end]) * (end - start)
		if area > maxArea {
			maxArea = area
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

