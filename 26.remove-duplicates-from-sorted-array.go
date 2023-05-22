package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	left, right := 0, 1

	for right < len(nums) {
		if nums[left] == nums[right] {
			nums = Remove(nums, right)
		} else {
			left = right
			right++
		}
	}

	return len(nums)
}

func Remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// @lc code=end
