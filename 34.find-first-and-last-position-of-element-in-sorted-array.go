package leetcode

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	a, b := -1, -1
	for left <= right && (a == -1 || b == -1) {
		if a == -1 && nums[left] == target {
			a = left
		} else if a == -1 {
			left++
		}

		if b == -1 && nums[right] == target {
			b = right
		} else if b == -1 {
			right--
		}
	}

	return []int{a, b}
}

// @lc code=end
