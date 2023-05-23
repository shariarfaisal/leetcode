package leetcode

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	combinations := [][]int{}

	backtrack(candidates, target, []int{}, 0, &combinations)

	return combinations
}

func backtrack(candidates []int, target int, current []int, start int, combinations *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*combinations = append(*combinations, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target {
			current = append(current, candidates[i])
			backtrack(candidates, target-candidates[i], current, i, combinations)
			current = current[:len(current)-1]
		}
	}
}

// @lc code=end
