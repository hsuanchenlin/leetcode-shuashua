/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (54.27%)
 * Total Accepted:    359K
 * Total Submissions: 661.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 * 
 */
 func permute(nums []int) [][]int {
	var tmp []int
	var ans [][]int
	backtrack(&ans, tmp, nums)
	return ans
}
func backtrack(list *[][]int, tmpList []int, nums []int){
	if len(tmpList) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, tmpList)
		*list = append(*list, tmp)
		return
	} else {
		for i:=0;i<len(nums);i++{
			if hasInSlice(tmpList, nums[i]) {
				continue
			}
			orilen := len(tmpList)
			tmpList = append(tmpList, nums[i])
			backtrack(list, tmpList, nums)	
			tmpList = tmpList[:orilen]
		}
	}	
}
func hasInSlice(list []int, target int) bool {
	for j:=0;j<len(list);j++{
		if list[j] == target{
			return true
		}
	}
	return false
}


