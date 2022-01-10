/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	i, j := 0, 0
	mid, prev := -1, -1

	for i+j <= length/2 {
		prev = mid
		if i < len(nums1) && (j == len(nums2) || nums1[i] <= nums2[j]) {
			mid = nums1[i]
			i++
		} else if j < len(nums2) && (i == len(nums1) || nums2[j] <= nums1[i]) {
			mid = nums2[j]
			j++
		}
	}

	if length&1 == 0 {
		return float64(prev+mid) / 2.0
	}

	return float64(mid)
}

// 方法二：移动下标，直到下标到达中位数
// 空间复杂度O(1)，时间复杂度O(m+n)

// @lc code=end
