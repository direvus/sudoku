package sudoku

import "fmt"

// CellRef represents a reference to a puzzle cell
type CellRef struct {
	row, col int
}

func (ref *CellRef) String() string {
	return fmt.Sprintf("R%vC%v", ref.row+1, ref.col+1)
}
