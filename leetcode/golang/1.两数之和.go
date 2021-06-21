/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}

	for i, n := range nums {
		if j, exist := numMap[target-n]; exist {
			return []int{i, j}
		} else {
			numMap[n] = i
		}
	}

	return nil
}

// 使用map，但是难点在于要考虑存在相等数字的情形，比如 3+3=6，
// 注意这里的限定条件只有一个答案

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// @lc code=end

