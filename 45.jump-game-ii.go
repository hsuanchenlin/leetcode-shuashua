/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (27.75%)
 * Total Accepted:    161.7K
 * Total Submissions: 582.4K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 * 
 * Each element in the array represents your maximum jump length at that
 * position.
 * 
 * Your goal is to reach the last index in the minimum number of jumps.
 * 
 * Example:
 * 
 * 
 * Input: [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2.
 * ⁠   Jump 1 step from index 0 to 1, then 3 steps to the last index.
 * 
 * Note:
 * 
 * You can assume that you can always reach the last index.
 * 
 */
func jump(nums []int) int {

	i, step, end := 0, 0, len(nums)-1

	var nextI, nextMax, maxi int
	for i < end {
		if i+nums[i] >= end{
			return step+1
		}
		nextI, nextMax = i+1, i+nums[i]
		for nextI <= nextMax{
			if nextI+nums[nextI] > maxi {
				maxi, i = nextI+nums[nextI], nextI
			}
			nextI++
		}
		step++
	}
	return step
}


