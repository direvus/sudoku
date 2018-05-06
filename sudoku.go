package sudoku

import (
	"bytes"
	"fmt"
)

// Size is the number of rows and columns of a sudoku puzzle.
const Size = 9

// Puzzle represents a 9×9 sudoku grid.
type Puzzle [Size][Size]byte

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
				puz[i][j] = ' '
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

// Validate a puzzle for correctness.
//
// A puzzle is incorrect if it contains the same glyph more than once on any
// line, any column, or in any of the nine 3×3 subgrids.
func (puz *Puzzle) Validate() error {
	return nil
}
