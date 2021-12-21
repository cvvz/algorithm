// https://leetcode-cn.com/problems/RQku0D/

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			if !isPalindrome(s[i:j]) && !isPalindrome(s[i+1:j+1]) {
				return false
			} else {
				return true
			}
		}
	}

	return true
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}