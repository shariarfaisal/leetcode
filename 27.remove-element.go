package leetcode

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	i := 0

	for i < len(nums) {
		if nums[i] == val {
			nums = remove(nums, i)
		} else {
			i++
		}
	}

	return len(nums)
}

func remove(n []int, i int) []int {
	return append(n[:i], n[i+1:]...)
}

// @lc code=end
