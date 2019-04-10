/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (40.78%)
 * Total Accepted:    210.1K
 * Total Submissions: 514.8K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 * 
 * Each number in candidates may only be used once in the combination.
 * 
 * Note:
 * 
 * 
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 * 
 */
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var tmp []int
	// ans := make([][]int, len(candidates))
	// tmp := make([]int, len(candidates))
	helper(candidates, target, tmp, ans)
	return ans
}
func helper(candidates []int, target int, tmp []int, ans [][]int){
	if target == 0 {
		ans = append(ans,tmp)
		return 
	}
	for i := 0;i < len(candidates); i++{
		if candidates[i] < target {
			newCandidate := []int {}
			newCandidate = append(candidates[:i], candidates[i+1:]...)
			helper(newCandidate, target-candidates[i], append(tmp, candidates[i]), ans)
		}
	}
}

