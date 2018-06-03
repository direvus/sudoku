package sudoku

import (
	"bytes"
	"fmt"
)

// SubSize is the number of rows and columns of a sudoku subgrid.
const SubSize = 3

// Size is the number of rows and columns of a sudoku puzzle.
const Size = SubSize * SubSize

// GridSize is the total number of cells in a sudoku puzzle grid.
const GridSize = Size * Size

// Unknown is the glyph that indicates a masked or unknown value.
const Unknown byte = ' '

// Null is the glyph that indicates an uninitialised value.
const Null byte = 0

// Glyphs contains all of the valid known glyphs.
var Glyphs = [Size]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

// Puzzle represents a 9×9 sudoku grid.
type Puzzle [GridSize]byte

// Known returns whether the given glyph indicates a known value.
func Known(glyph byte) bool {
	return glyph >= '1' && glyph <= '9'
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
	i := 0
	for r, line := range lines {
		// Expecting one byte separating each puzzle glyph.
		length := (Size * 2) - 1
		if len(line) != length {
			return fmt.Errorf("malformed input on line %v: expected %v bytes, got %v", r+1, len(line), length)
		}
		for c := 0; c < Size; c++ {
			glyph := line[c*2]
			if Known(glyph) {
				puz[i] = glyph
			} else if glyph == '_' || glyph == ' ' {
				puz[i] = Unknown
			} else {
				return fmt.Errorf("malformed input on line %v: expected underscore, space or digit 1-9 in column %v, got %q", r+1, c+1, glyph)
			}
			i++
		}
	}
	return nil
}

// coordsToIndex returns the grid index for a given row and column index.
func coordsToIndex(row, col int) int {
	return row * Size + col
}

// indexToCoords returns the row and column index for a given grid index.
func indexToCoords(index int) (row, col int) {
	row = index / Size
	col = index % Size
	return
}

// cellRefToIndex returns the grid index for a given CellRef.
func cellRefToIndex(cell CellRef) int {
	return coordsToIndex(cell.row, cell.col)
}

// indexToCellRef returns the CellRef for a given grid index.
func indexToCellRef(index int) CellRef {
	return CellRef{index / Size, index % Size}
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
	i := index * Size
	return puz[i:i+Size]
}

// Column returns one column from a puzzle as a slice of bytes.
//
// Columns are indexed from left to right, beginning with zero.
func (puz *Puzzle) Column(index int) []byte {
	var col [Size]byte
	for r := 0; r < Size; r++ {
		col[r] = puz[r*Size+index]
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
	return ((r / SubSize) * SubSize) + (c / SubSize)
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
	c := (index / SubSize) * SubSize * Size + (index % SubSize * SubSize)
	var subgrid [Size]byte
	for i := 0; i < Size; i++ {
		if i > 0 && i % 3 == 0 {
			c += Size - SubSize
		}
		subgrid[i] = puz[c]
		c++
	}
	return subgrid[:]
}

// Equal returns whether two puzzles contain the same bytes.
func (a *Puzzle) Equal(b Puzzle) bool {
	return bytes.Equal(a[:], b[:])
}

// NumUnknowns returns the number of unknown cells in the puzzle.
//
// Null bytes count as unknown for this method.
func (puz *Puzzle) NumUnknowns() (count int) {
	for i := 0; i < GridSize; i++ {
		if !Known(puz[i]) {
			count++
		}
	}
	return
}

// NextUnknown returns the location of the next unknown cell.
//
// The search is performed in top-down, left-right order beginning at the given
// location, and returns the location of the first unknown cell found.
//
// If no unknown cells are found in the search, 'found' will be returned as
// false.
func (puz *Puzzle) NextUnknown(r, c int) (row, column int, found bool) {
	for i := coordsToIndex(r, c); i < GridSize; i++ {
		if !Known(puz[i]) {
			row, column = indexToCoords(i)
			found = true
			return
		}
	}
	return
}

// FindUnknown returns the location of the nearest unknown cell.
//
// The search is performed in top-down, left-right order beginning at the given
// location, wrapping back to R1C1 if it does not begin there, and returns the
// location of the first unknown cell found.
//
// If no unknown cells are found in the search, 'found' will be returned as
// false.
func (puz *Puzzle) FindUnknown(r, c int) (row, column int, found bool) {
	index := coordsToIndex(r, c)
	for i, v := range puz[index:] {
		if !Known(v) {
			row, column = indexToCoords(index+i)
			found = true
			return
		}
	}
	if index != 0 {
		for i, v := range puz[:index] {
			if !Known(v) {
				row, column = indexToCoords(i)
				found = true
				return
			}
		}
	}
	return
}

// Unknowns returns a slice of all unknown CellRefs in the puzzle.
func (puz *Puzzle) Unknowns() (refs []CellRef) {
	for i := 0; i < GridSize; i++ {
		if !Known(puz[i]) {
			refs = append(refs, indexToCellRef(i))
		}
	}
	return
}

// Knowns returns a slice of all known CellRefs in the puzzle.
func (puz *Puzzle) Knowns() (refs []CellRef) {
	for i := 0; i < GridSize; i++ {
		if Known(puz[i]) {
			refs = append(refs, indexToCellRef(i))
		}
	}
	return
}

// GetCell returns the value of the cell at the given CellRef.
func (puz *Puzzle) GetCell(cell CellRef) byte {
	return puz[cellRefToIndex(cell)]
}

// SetCell sets the value of the cell at the given CellRef.
func (puz *Puzzle) SetCell(cell CellRef, v byte) {
	puz[cellRefToIndex(cell)] = v
}

// Merge copies bytes from 'source' into 'dest'.
//
// Null (0x00) bytes in the source are disregarded.
func (dest *Puzzle) Merge(source Puzzle) {
	for i := 0; i < GridSize; i++ {
		if source[i] != 0 {
			dest[i] = source[i]
		}
	}
}

// ApplyMask returns a new puzzle with a boolean mask applied.
//
// For each cell in the source puzzle, check the corresponding cell in the
// given mask object.  If the cell is true in the mask, then the cell in the
// output puzzle has the same value as the cell in the source puzzle.  If the
// cell is false in the mask, the cell in the output puzzle is Unknown.
func (source *Puzzle) ApplyMask(mask *Mask) (puzzle Puzzle) {
	for i := 0; i < GridSize; i++ {
		if mask[i] {
			puzzle[i] = source[i]
		} else {
			puzzle[i] = Unknown
		}
	}
	return
}

// GetMask returns a mask corresponding to known values in the puzzle.
//
// For every known value in the puzzle, the corresponding cell in the mask will
// be true.  For every unknown value, the corresponding cell in the mask will
// be false.
func (puz *Puzzle) GetMask() (mask Mask) {
	for i := 0; i < GridSize; i++ {
		mask[i] = Known(puz[i])
	}
	return
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
	for i := 0; i < GridSize; i++ {
		glyph := puz[i]
		if !Known(glyph) {
			buf.WriteByte('_')
		} else {
			buf.WriteByte(puz[i])
		}
		if i % Size == Size-1 {
			buf.WriteByte('\n')
		} else {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

// Clear sets all bytes of the puzzle to Null.
func (puz *Puzzle) Clear() {
	for i := 0; i < GridSize; i++ {
		puz[i] = Null
	}
	return
}
