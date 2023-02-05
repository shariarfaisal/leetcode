package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var sign int
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	} else {
		sign = 1
	}
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		res = res*10 + int(s[i]-'0')
		if res > 2147483647 {
			break
		}
	}
	res = res * sign
	if res > 2147483647 {
		return 2147483647
	}
	if res < -2147483648 {
		return -2147483648
	}
	return res
}

// @lc code=end
