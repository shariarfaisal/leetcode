package leetcode

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	var res int
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res > 1<<31-1 {
		return 0
	}

	return sign * res
}

// @lc code=end
