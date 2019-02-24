package leetcode

import (
	"errors"
	"math"
)

// Point in 2D.
type Point struct {
	X int
	Y int
}

// Board for playing chess.
type Board struct {
	Size   int
	Queens []Point
}

// NewEmptyBoard creates an empty board of the given size.
func NewEmptyBoard(size int) *Board {
	return &Board{
		Size: size,
	}
}

// Place a queen at the given position.
func (b *Board) Place(x, y int) {
	b.Queens = append(b.Queens, Point{X: x, Y: y})
}

// CanPlace returns an error if a queen can't be placed at the given position.
func (b *Board) CanPlace(x, y int) error {
	for _, q := range b.Queens {
		if q.X == x {
			return errors.New("can't place on same row")
		}

		if q.Y == y {
			return errors.New("can't place on same column")
		}

		if math.Abs(float64(q.X-x)) == math.Abs(float64(q.Y-y)) {
			return errors.New("can't place on diagonal")
		}
	}

	return nil
}

// Copy a board (deep copy).
func (b *Board) Copy() *Board {
	queens := make([]Point, len(b.Queens))
	for i := 0; i < len(queens); i++ {
		queens[i] = b.Queens[i]
	}

	return &Board{
		Size:   b.Size,
		Queens: queens,
	}
}

// Complete returns true if the board contains one queen per row.
func (b *Board) Complete() bool {
	return len(b.Queens) == b.Size
}

// Export the board to the solution format.
func (b *Board) Export() []string {
	bb := make([][]byte, b.Size)
	for i := 0; i < b.Size; i++ {
		bb[i] = make([]byte, b.Size)
		for j := 0; j < b.Size; j++ {
			bb[i][j] = '.'
		}
	}

	for _, q := range b.Queens {
		bb[q.X][q.Y] = 'Q'
	}

	sol := make([]string, b.Size)
	for i := 0; i < b.Size; i++ {
		sol[i] = string(bb[i])
	}

	return sol
}

func solveNQueens(size int) [][]string {
	boards := solveNQueensRec(NewEmptyBoard(size))
	var solutions [][]string
	for _, b := range boards {
		solutions = append(solutions, b.Export())
	}

	return solutions
}

func solveNQueensRec(board *Board) []*Board {
	if board.Complete() {
		return []*Board{board}
	}

	var boards []*Board

	row := len(board.Queens)
	for col := 0; col < board.Size; col++ {
		err := board.CanPlace(row, col)
		if err != nil {
			continue
		}

		candidate := board.Copy()
		candidate.Place(row, col)

		boards = append(boards, solveNQueensRec(candidate)...)
	}

	return boards
}
