// https://leetcode-cn.com/problems/XltzEq/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1

	for i < j {
		if !((s[i] <= 'z' && s[i] >= 'a') || (s[i] <= '9' && s[i] >= '0')) {
			i++
			continue
		}
		if !((s[j] <= 'z' && s[j] >= 'a') || (s[j] <= '9' && s[j] >= '0')) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}