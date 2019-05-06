/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 *
 * https://leetcode.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (37.79%)
 * Likes:    1599
 * Dislikes: 94
 * Total Accepted:    153.2K
 * Total Submissions: 404.3K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * Given an array nums, there is a sliding window of size k which is moving
 * from the very left of the array to the very right. You can only see the k
 * numbers in the window. Each time the sliding window moves right by one
 * position. Return the max sliding window.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
 * Output: [3,3,5,5,6,7] 
 * Explanation: 
 * 
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty
 * array.
 * 
 * Follow up:
 * Could you solve it in linear time?
 */
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int {}
	}
	tmp := make([]int, k)
	copy(tmp, nums[:k])
	tmpMaxIdx := maxIdx(tmp) 
	res :=  make([]int, len(nums)-k+1) 
	res[0] = tmp[tmpMaxIdx]
	for i, v := range nums[k:] {
		if v > tmp[tmpMaxIdx] {
			tmp = tmp[len(tmp)-1:]
			tmp[0] = v
			res[i+1] = v
			tmpMaxIdx = 0
		} else {
			if len(tmp) == k {
				tmp = tmp[1:]
			}
			tmp = append(tmp, v)
			if tmpMaxIdx == 0 {
				tmpMaxIdx = maxIdx(tmp)
			} else {
				tmpMaxIdx--
			}
			res[i+1] = tmp[tmpMaxIdx]
		}
	} 
	return res
}

func maxIdx(nums []int) int {
	idx := 0
	tmpMax := nums[0]
	for i, v := range nums {
		if v > tmpMax{
			idx = i
			tmpMax = v
		}
	}
	return idx
}