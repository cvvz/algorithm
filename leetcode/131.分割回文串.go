/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	ans := [][]string{}
	subAns := []string{
		string(s[0]),
	}
	n := len(s)

	var backtracking func(int)
	backtracking = func(i int) {
		if i == n {
			for _, str := range subAns {
				if !isLegal(str) {
					return
				}
			}
			ans = append(ans, append([]string{}, subAns...))
			return
		}

		letter := string(s[i])

		// 单独拿出来
		subAns = append(subAns, letter)
		backtracking(i + 1)
		subAns = subAns[:len(subAns)-1]

		// 或者拿出来和前面的子串合并
		prev := subAns[len(subAns)-1]
		subAns[len(subAns)-1] = prev + letter
		backtracking(i + 1)
		subAns[len(subAns)-1] = prev
	}

	backtracking(1)

	return ans
}

func isLegal(str string) bool {
	n := len(str)

	start, end := 0, n-1
	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// @lc code=end

