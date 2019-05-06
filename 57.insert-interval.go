/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (31.03%)
 * Total Accepted:    172K
 * Total Submissions: 554.1K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * Given a set of non-overlapping intervals, insert a new interval into the
 * intervals (merge if necessary).
 * 
 * You may assume that the intervals were initially sorted according to their
 * start times.
 * 
 * Example 1:
 * 
 * 
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
    if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int

	cur := intervals[0]
	for i:=1;i<len(intervals);i++ {
		if(intervals[i][0] <= cur[1]){
			cur[1] = max(cur[1], intervals[i][1])
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	// fmt.Println(intervals)
	return res
}

func max(a, b int) int {
	if(a > b){
		return a
	} 
	return b
}


