package sudoku

// guessCount counts the number of solutions to a puzzle by guesswork.
//
// Starting with the given cell, guessCount tries each glyph in turn, and if it
// finds a valid glyph, recurses on to the next unknown cell and repeats the
// process.
//
// The channel receives the number of valid solutions found for this cell and
// all subsequent cells.  If at any point the method detects that more than one
// solution has been found, it immediately writes to the channel and returns.
func (puz *Puzzle) guessCount(r, c int, ch chan int) {
	subgrid := CellSubGrid(r, c)
	index := coordsToIndex(r, c)
	count := 0
	orig := puz[index]
	for _, glyph := range Glyphs {
		if (puz.glyphInRow(glyph, r, c) ||
				puz.glyphInColumn(glyph, c, r) ||
				puz.glyphInSubGrid(glyph, subgrid, r, c)) {
			continue
		}
		puz[index] = glyph
		if puz.Validate() == nil {
			nr, nc, found := puz.NextUnknown(r, c)
			if found {
				// So far so good, recurse to the next cell.
				nch := make(chan int)
				go puz.guessCount(nr, nc, nch)
				count += <-nch
			} else {
				count++
			}
			if count > 1 {
				ch <- count
				return
			}
		}
	}
	puz[index] = orig
	ch <- count
}

// NumSolutions returns the number of solutions to a puzzle.
//
// It first solves all cells which can be determined by simple logical
// elimination, and then begins guesswork with backtracking to discover all
// possible solutions to the remaining unknown cells.  If at any point it
// becomes clear that multiple solutions exist, the method returns an integer
// greater than one.
//
// Otherwise, it continues until all possibilities have been exhausted and
// returns one if a solution has been found, zero if it has not.
//
// This method will modify the contents of 'puz', so make a copy to pass in
// here, if you want to keep the original.
func (puz *Puzzle) NumSolutions() int {
	puz.SolveEasy()
	r, c, found := puz.NextUnknown(0, 0)
	if !found {
		return 1
	}
	ch := make(chan int)
	go puz.guessCount(r, c, ch)
	return <-ch
}
