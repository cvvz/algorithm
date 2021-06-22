/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	symbolValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	n := len(s)
	ans := 0
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans

}

// 从左到右依次遍历，只要左边的数字小于右边，则将左边的数字取反

// 时间复杂度 O(n)
// 空间复杂度 O(1) // 注意：如果先把s string转换成[]byte，则空间复杂度为O(n)

// 备忘录📕：
// 1. go中单引号默认是rune类型，但是如果单字节能表示（ASCII）的话，那么也可以是byte类型
// 2. 使用for range遍历string类型时，默认返回下标，而不是rune
// 3. string类型通过下标访问到的是byte而不是rune
// @lc code=end

