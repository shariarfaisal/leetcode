package leetcode

import "math"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1

	maxArea := 0

	for left < right {
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// @lc code=end
