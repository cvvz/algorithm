/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start

// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2) == 0 {
		return nil
	}
	stack := []int{nums2[0]}
	greater := make(map[int]int, 0)
	var ret []int
	for _, val := range nums2[1:] {
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] > val {
				break
			}
			greater[stack[i]] = val
			stack = stack[:i]
		}
		stack = append(stack, val)
	}
	for _, val := range nums1 {
		if v, ok := greater[val]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}

// 为什么有两重循环，时间复杂度还是O(n)?
// 因为入栈语句执行了O(n)次，出栈语句最多也只会执行O(n)次
// @lc code=end

