package main

import "strconv"

const n = 3
const N = 9

type Solver struct {
	rows         [][]int
	columns      [][]int
	boxes        [][]int
	board        [][]byte
	sudokuSolved bool
}

func buildSolver() *Solver {
	rows := [][]int{}
	columns := [][]int{}
	boxes := [][]int{}

	for i := 0; i < 9; i++ {
		rows = append(rows, []int{})
		columns = append(columns, []int{})
		boxes = append(boxes, []int{})
	}

	solver := &Solver{rows: rows, columns: columns, boxes: boxes}
	return solver
}

func (s *Solver) couldPlace(d, row, col int) bool {
	/*
	   Check if one could place a number d in (row, col) cell
	*/
	idx := (row/n)*n + col/n
	return s.rows[row][d]+s.columns[col][d]+s.boxes[idx][d] == 0
}

func (s *Solver) placeNumber(d, row, col int) {
	/*
	   Place a number d in (row, col) cell
	*/
	idx := (row/n)*n + col/n

	s.rows[row][d]++
	s.columns[col][d]++
	s.boxes[idx][d]++
	s.board[row][col] = (byte)(d + '0')
}

func (s *Solver) removeNumber(d, row, col int) {
	/*
	   Remove a number which didn't lead to a solution
	*/
	idx := (row/n)*n + col/n

	s.rows[row][d]--
	s.columns[col][d]--
	s.boxes[idx][d]--
	s.board[row][col] = '.'
}

func (s *Solver) placeNextNumber(row, col int) {
	/*
	   Call backtrack function in recursion
	   to continue to place numbers
	   till the moment we have a solution
	*/
	// if we're in the last cell
	// that means we have the solution
	if (col == N-1) && (row == N-1) {
		s.sudokuSolved = true
	} else {
		if col == N-1 {
			s.backtrace(row+1, 0)
		} else {
			s.backtrace(row, col+1)
		}
	}
}

func (s *Solver) backtrace(row, col int) {
	if s.board[row][col] == '.' {
		for d := 1; d < 10; d++ {
			s.placeNumber(d, row, col)
			s.placeNextNumber(row, col)
			if !s.sudokuSolved {
				s.removeNumber(d, row, col)
			}
		}
	} else {
		s.placeNextNumber(row, col)
	}
}

func solveSudoku(board [][]byte) {
	solver := buildSolver()
	solver.board = board

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] != '.' {
				d, _ := strconv.Atoi(string(board[i][j]))
				solver.placeNumber(d, i, j)
			}
		}
	}
	solver.backtrace(0, 0)
}
