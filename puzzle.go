package sudoku

import (
	"bytes"
	"fmt"
)

// Size is the number of rows and columns of a sudoku puzzle.
const Size = 9

// SubSize is the number of rows and columns of a sudoku subgrid.
const SubSize = 3

// Unknown is the glyph that indicates a masked or unknown value.
const Unknown byte = ' '

// Null is the glyph that indicates an uninitialised value.
const Null byte = 0

// Glyphs contains all of the valid known glyphs.
var Glyphs = [Size]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

// Puzzle represents a 9×9 sudoku grid.
type Puzzle [Size][Size]byte

// Known returns whether the given glyph indicates a known value.
//
// That is, it is neither Unknown, nor Null.
func Known(glyph byte) bool {
	return glyph != Unknown && glyph != Null
}

// Read in a puzzle definition from a slice of bytes.
//
// The input format should contain one line for each puzzle row, with
// lines containing puzzle glyphs separated by a single byte.  Unknown
// (masked) values should be indicated by the underscore character 0x5f.
//
// E.g., the following is valid puzzle input:
// 1 _ 3 _ _ 6 _ 8 _
// _ 5 _ _ 8 _ 1 2 _
// 7 _ 9 1 _ 3 _ 5 6
// _ 3 _ _ 6 7 _ 9 _
// 5 _ 7 8 _ _ _ 3 _
// 8 _ 1 _ 3 _ 5 _ 7
// _ 4 _ _ 7 8 _ 1 _
// 6 _ 8 _ _ 2 _ 4 _
// _ 1 2 _ 4 5 _ 7 8
func (puz *Puzzle) Read(input []byte) error {
	ending := []byte("\n")
	lines := bytes.Split(bytes.TrimSpace(input), ending)
	if len(lines) != Size {
		return fmt.Errorf("malformed input: expected %v lines, got %v", Size, len(lines))
	}
	for i, line := range lines {
		// Expecting one byte separating each puzzle glyph.
		length := (Size * 2) - 1
		if len(line) != length {
			return fmt.Errorf("malformed input on line %v: expected %v bytes, got %v", i+1, len(line), length)
		}
		for j := 0; j < Size; j++ {
			glyph := line[j*2]
			if glyph == '_' {
				// Masked value, set cell to space for "unknown".
				puz[i][j] = Unknown
			} else if glyph >= '1' && glyph <= '9' {
				// Known value.
				puz[i][j] = glyph
			} else {
				return fmt.Errorf("malformed input on line %v: expected underscore or digit 1-9 in column %v, got %v", i+1, j, glyph)
			}
		}
	}
	return nil
}

// findDuplicate searches the argument for duplicate glyphs, and returns the
// first glyph which occurs more than once.  It returns the null byte 0x00 if
// no duplicates exist.  Duplicates of unknown bytes are disregarded.
func findDuplicate(input []byte) byte {
	// Use an empty struct mapping as a poor man's "set" type.
	var seen map[byte]struct{} = make(map[byte]struct{})
	for _, v := range input {
		if !Known(v) {
			continue
		}
		if _, ok := seen[v]; ok {
			return v
		}
		seen[v] = struct{}{}
	}
	return 0
}

// Row returns one row from a puzzle as a slice of bytes.
//
// Rows are indexed from top to bottom, beginning with zero.
func (puz *Puzzle) Row(index int) []byte {
	return puz[index][:]
}

// Column returns one column from a puzzle as a slice of bytes.
//
// Columns are indexed from left to right, beginning with zero.
func (puz *Puzzle) Column(index int) []byte {
	var col [Size]byte
	for r := 0; r < Size; r++ {
		col[r] = puz[r][index]
	}
	return col[:]
}

// CellSubGrid returns the index of the subgrid that the given cell is in.
//
// Subgrids are indexed in left to right, top to bottom order, beginning with
// zero:
//
// 0 1 2
// 3 4 5
// 6 7 8
func CellSubGrid(r, c int) int {
	return ((r / 3) * 3) + (c / 3)
}

// SubGrid returns one subgrid from a puzzle as a slice of bytes.
//
// Subgrids are indexed in left to right, top to bottom order, beginning with
// zero.  The returned slice contains bytes from the subgrid in the same order:
//
// 0 1 2
// 3 4 5
// 6 7 8
func (puz *Puzzle) SubGrid(index int) []byte {
	r := (index / SubSize) * SubSize
	c := (index % SubSize) * SubSize
	subgrid := []byte{}
	for i := 0; i < SubSize; i++ {
		for j := 0; j < SubSize; j++ {
			subgrid = append(subgrid, puz[r+i][c+j])
		}
	}
	return subgrid
}

// Equal returns whether two puzzles contain the same bytes.
func (a *Puzzle) Equal(b Puzzle) bool {
	for i := 0; i < Size; i++ {
		if !bytes.Equal(a[i][:], b[i][:]) {
			return false
		}
	}
	return true
}

// NumUnknowns returns the number of unknown cells in the puzzle.
//
// Null bytes count as unknown for this method.
func (puz *Puzzle) NumUnknowns() (count int) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if !Known(puz[i][j]) {
				count++
			}
		}
	}
	return
}

// NextUnknown returns the location of the next unknown cell.
//
// The search is performed in top-down, left-right order, and retuns the
// location of the first unknown cell found.
//
// If no unknown cells exist, 'found' will be returned as false.
func (puz *Puzzle) NextUnknown(r, c int) (row, column int, found bool) {
	for row = r; row < Size; row++ {
		for column = c; column < Size; column++ {
			if !Known(puz[row][column]) {
				found = true
				return
			}
		}
		c = 0
	}
	return
}

// Merge copies bytes from 'source' into 'dest'.
//
// Null (0x00) bytes in the source are disregarded.
func (dest *Puzzle) Merge(source Puzzle) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if source[i][j] != 0 {
				dest[i][j] = source[i][j]
			}
		}
	}
}

// Validate a puzzle for correctness.
//
// A puzzle is incorrect if it contains the same glyph more than once on any
// line, any column, or in any of the nine 3×3 subgrids.
func (puz *Puzzle) Validate() error {
	// Rows
	for i := 0; i < Size; i++ {
		dup := findDuplicate(puz.Row(i))
		if dup != 0 {
			return fmt.Errorf("invalid puzzle: duplicate %q in row %v", dup, i+1)
		}
	}
	// Columns
	for i := 0; i < Size; i++ {
		dup := findDuplicate(puz.Column(i))
		if dup != 0 {
			return fmt.Errorf("invalid puzzle: duplicate %q in column %v", dup, i+1)
		}
	}
	// Subgrids
	for i := 0; i < Size; i++ {
		dup := findDuplicate(puz.SubGrid(i))
		if dup != 0 {
			return fmt.Errorf("invalid puzzle: duplicate %q in subgrid %v", dup, i+1)
		}
	}
	return nil
}

// String returns a formatted representation of a puzzle.
//
// Rows are each terminated by a newline (0x0a), while glyphs within a row are
// separated by one space (0x20).  Nulls and unknowns are represented by
// underscore (0x5f).
//
// This format can be consumed by the Read() method.
func (puz *Puzzle) String() string {
	var buf bytes.Buffer
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			glyph := puz[i][j]
			if !Known(glyph) {
				buf.WriteByte('_')
			} else {
				buf.WriteByte(puz[i][j])
			}
			if j < Size-1 {
				buf.WriteByte(' ')
			} else {
				buf.WriteByte('\n')
			}
		}
	}
	return buf.String()
}

// Clear sets all bytes of the puzzle to Null.
func (puz *Puzzle) Clear() {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			puz[i][j] = Null
		}
	}
	return
}
