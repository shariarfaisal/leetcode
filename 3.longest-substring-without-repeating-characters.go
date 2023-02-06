package leetcode

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]int)
		m[s[i]] = 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = 1
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}

// @lc code=end
