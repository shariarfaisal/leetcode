package leetcode

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {

	// init
	// 1. init tMap
	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	// 2. init sMap
	sMap := make(map[byte]int)
	// 3. init left, right, minLen, minLeft
	left, right, minLen, minLeft := 0, 0, len(s)+1, 0
	// 4. init valid
	valid := 0

	// loop
	for right < len(s) {
		// 1. add right
		c := s[right]
		right++
		// 2. update sMap
		if _, ok := tMap[c]; ok {
			sMap[c]++
			if sMap[c] == tMap[c] {
				valid++
			}
		}
		// 3. shrink left
		for valid == len(tMap) {
			// 3.1 update minLen, minLeft
			if right-left < minLen {
				minLen = right - left
				minLeft = left
			}
			// 3.2 remove left
			d := s[left]
			left++
			// 3.3 update sMap
			if _, ok := tMap[d]; ok {
				if sMap[d] == tMap[d] {
					valid--
				}
				sMap[d]--
			}
		}
	}

	// return
	if minLen == len(s)+1 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// @lc code=end
