/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i,j,k:=m-1,n-1,m+n-1;i >=0 || j>=0;k-- {
		if i <0 {
			nums1[k] = nums2[j]
			j--
			continue
		}
		if j <0{
			nums1[k] = nums1[i]
			i--
			continue
		} 

		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}


// 逆向双指针
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
// @lc code=end

