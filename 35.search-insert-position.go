package leetcode

import "math"

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end
