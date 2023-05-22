package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	pivot := findPivot(nums, 0, len(nums)-1)
	if pivot == -1 {
		return binarySearch(nums, target, 0, len(nums)-1)
	} else if nums[pivot] == target {
		return pivot
	} else if nums[0] <= target {
		return binarySearch(nums, target, 0, pivot-1)
	} else {
		return binarySearch(nums, target, pivot+1, len(nums)-1)
	}

}

func binarySearch(nums []int, target int, left int, right int) int {
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

	return -1
}

func findPivot(nums []int, left int, right int) int {
	if nums[left] <= nums[right] {
		return -1
	}

	for left <= right {
		mid := int(math.Floor(float64(left+right) / 2))
		next := (mid + 1)

		if nums[mid] > nums[next] {
			return next
		} else if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// @lc code=end
