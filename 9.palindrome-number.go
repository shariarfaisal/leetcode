package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var y int
	for x > y {
		y = y*10 + x%10
		x = x / 10
		println(x, y)
	}
	return x == y || x == y/10
}

// @lc code=end
