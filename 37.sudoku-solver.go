/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (36.12%)
 * Total Accepted:    122.6K
 * Total Submissions: 339.2K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 * 
 * A sudoku solution must satisfy all of the following rules:
 * 
 * 
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 * 
 * 
 * Empty cells are indicated by the character '.'.
 * 
 * 
 * A sudoku puzzle...
 * 
 * 
 * ...and its solution numbers marked in red.
 * 
 * Note:
 * 
 * 
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 * 
 * 
 */
func solveSudoku(board [][]byte)  {
	cnt := 0
	
	solve(board, cnt)
}
func solve(board [][]byte, cnt int) bool{
	if cnt == 81 {
		return true
	}
	// for i:=0;i<9;i++{
	// 	if !isRowValid() && !isColValid() && !isSqValid() {
	// 		return 
	// 	}	
	// }
	// solve()
	
	r, c := cnt/9, cnt%9
	if board[r][c] != '.'{
		return solve(board, cnt+1)
	}
	bi, bj := r/3*3, c/3*3
	isValid := func(b byte) bool{
		for n:=0;n<9;n++{
			if board[r][n] == b || board[n][c] == b || board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}
	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if solve(board, cnt+1) {
				return true
			}
		}
	} 
	board[r][c] ='.'
	return false
}

