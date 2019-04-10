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
type alphabetical [][]int

func (s alphabetical) Len() int {
	return len(s)
}

func (s alphabetical) Swap(i, j int){
	 s[i], s[j] = s[j], s[i]
}
func (s alphabetical) Less(i, j int) bool {
	if len(s[i]) != len(s[j]){
    	return len(s[i]) < len(s[j])
	} else {
		for k:=0;k<len(s[i]);k++{
			if s[i][k] != s[j][k]{
				return s[i][j] < s[j][k] 
			}
		}
		return len(s[i]) < len(s[j]) 
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans :=  [][]int {}
	tmp := []int{}
	helper(candidates, target, tmp, &ans)
	// sort.Sort(alphabetical(ans))
	return ans
}
func helper(candidates []int, target int, tmp []int, ans *[][]int){
	if target == 0 {
		*ans = append(*ans,tmp)
		return 
	}
	if candidates[0] > target || len(candidates) == 0 {
		return 
	}
	helper(candidates[1:], target-candidates[0], append(tmp, candidates[0]), ans)
	helper(next(candidates), target, tmp, ans)
}


func next(candidates []int) []int {
	for i:=1;i<len(candidates);i++{
		if candidates[i] != candidates[0] {
			return candidates[i:] 
		}
	}
	return candidates
}

