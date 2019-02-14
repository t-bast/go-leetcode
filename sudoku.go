package leetcode

func solveSudoku(board [][]byte) {
	if !solveSudokuRec(board) {
		panic("sudoku can't be solved")
	}
}

func solveSudokuRec(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// If we can't find any valid candidates, we messed up.
				// The caller needs to its attempted change.
				cand := candidates(board, i, j)
				if len(cand) == 0 {
					return false
				}

				for _, candidate := range cand {
					board[i][j] = candidate
					rec := solveSudokuRec(board)
					if rec {
						return true
					}

					// Backtrack. Recursion does the magic.
					board[i][j] = '.'
				}

				return false
			}
		}
	}

	return true
}

func candidates(board [][]byte, i, j int) []byte {
	cand := map[byte]struct{}{
		'1': struct{}{},
		'2': struct{}{},
		'3': struct{}{},
		'4': struct{}{},
		'5': struct{}{},
		'6': struct{}{},
		'7': struct{}{},
		'8': struct{}{},
		'9': struct{}{},
	}

	for n := 0; n < 9; n++ {
		delete(cand, board[i][n])                     // line
		delete(cand, board[n][j])                     // column
		delete(cand, board[3*(i/3)+n%3][3*(j/3)+n/3]) // square
	}

	var res []byte
	for v := range cand {
		res = append(res, v)
	}

	return res
}

func verifySudoku(board [][]byte) bool {
	// Reject unfinished boards.
	for _, line := range board {
		for _, val := range line {
			if val < '1' || '9' < val {
				return false
			}
		}
	}

	// Verify lines.
	for i := 0; i < 9; i++ {
		sum := byte(0)
		seen := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			val := board[i][j]
			_, ok := seen[val]
			if ok {
				return false
			}

			seen[val] = struct{}{}
			sum += val - '0'
		}

		if sum != byte(45) {
			return false
		}
	}

	// Verify columns.
	for i := 0; i < 9; i++ {
		sum := byte(0)
		seen := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			val := board[j][i]
			_, ok := seen[val]
			if ok {
				return false
			}

			seen[val] = struct{}{}
			sum += val - '0'
		}

		if sum != byte(45) {
			return false
		}
	}

	// Verify squares.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := byte(0)
			seen := make(map[byte]struct{})

			for k := 0; k < 9; k++ {
				x := 3*i + k%3
				y := 3*j + k/3
				val := board[x][y]
				_, ok := seen[val]
				if ok {
					return false
				}

				seen[val] = struct{}{}
				sum += val - '0'
			}

			if sum != byte(45) {
				return false
			}
		}
	}

	return true
}
