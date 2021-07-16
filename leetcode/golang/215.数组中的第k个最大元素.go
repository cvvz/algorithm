/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// 快速排序

}

// 冒泡排序
// func findKthLargest(nums []int, k int) int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums)-i-1; j++ {
// 			if nums[j] > nums[j+1] {
// 				nums[j+1], nums[j] = nums[j], nums[j+1]
// 			}
// 		}
// 		if i == k-1 {
// 			return nums[len(nums)-1-i]
// 		}
// 	}

// 	return 0
// }

// @lc code=end

