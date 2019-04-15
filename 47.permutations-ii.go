/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (39.85%)
 * Total Accepted:    232.7K
 * Total Submissions: 584K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 * 
 */
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var tmp []int
	var ans [][]int
	taken := make([]bool, len(nums))
	backtrack(&ans, tmp, nums, taken)
	return ans
}
func backtrack(list *[][]int, tmpList []int, nums []int, taken []bool){
	if len(tmpList) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, tmpList)
		*list = append(*list, tmp)
		return
	} else {
		for i:=0;i<len(nums);i++ {
			if taken[i]{
				continue
			}
			if i>0 && nums[i-1] == nums[i] && !taken[i-1] {
				continue
			}
			orilen := len(tmpList)
			tmpList = append(tmpList, nums[i])
			taken[i] = true
			backtrack(list, tmpList, nums, taken)	
			tmpList = tmpList[:orilen]
			taken[i] = false
		}
	}	
}
func hasInSlice(list []int, target int, cnt int) bool {
	acc := 1
	for j:=0;j<len(list);j++{
		if list[j] == target {
			if acc >= cnt {
				return true
			}  
			acc++
		}
	}
	return false
}
