/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (42.54%)
 * Total Accepted:    274.2K
 * Total Submissions: 644.5K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 * 
 * 
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 * 
 * Example:
 * 
 * 
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * 
 */
func trap(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	left := make([]int, length)
	right := make([]int, length)
	
	left[0] = height[0]
	right[length-1] = height[length-1]
	for i:=1;i<length;i++{
		left[i] = max(left[i-1], height[i])
		right[length-i-1] = max(right[length-i], height[length-i-1])
	}
	ans := 0
	for i:=0;i<length;i++ {
		ans += min(left[i], right[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
