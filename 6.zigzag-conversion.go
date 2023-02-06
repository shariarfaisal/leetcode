package leetcode

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res string
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += 2 * (numRows - 1) {
			res += string(s[j+i])
			if i != 0 && i != numRows-1 && j+2*(numRows-1)-i < len(s) {
				res += string(s[j+2*(numRows-1)-i])
			}
		}
	}
	return res
}

// @lc code=end
