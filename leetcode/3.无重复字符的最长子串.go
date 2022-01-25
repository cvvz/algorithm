/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	sLen := len(s)
	maxLen, curLen := 0, 0
	hashTable := map[byte]int{}

	for left, right := 0, 0; right < sLen; right++ {
		if index, exist := hashTable[s[right]]; !exist {
			curLen++
		} else {
			for ; left <= index; left++ {
				delete(hashTable, s[left])
			}
			curLen = right - left + 1
		}
		hashTable[s[right]] = right
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
// @lc code=end

