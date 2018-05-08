package sudoku

// Candidates returns all of the candidate glyphs for a given puzzle cell.
func (puz *Puzzle) Candidates(r, c int) (result []byte) {
	candidates := map[byte]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true}
	// Eliminate glyphs in same row.
	for i := 0; i < Size; i++ {
		if i != c {
			glyph := puz[r][i]
			if glyph != Unknown && glyph != 0 {
				candidates[glyph] = false
			}
		}
	}
	// Eliminate glyphs in same column.
	for i := 0; i < Size; i++ {
		if i != r {
			glyph := puz[i][c]
			if glyph != Unknown && glyph != 0 {
				candidates[glyph] = false
			}
		}
	}
	// Eliminate glyphs in same subgrid.
	sr := (r / SubSize) * SubSize
	sc := (c / SubSize) * SubSize
	for i := sr; i < sr+SubSize; i++ {
		for j := sc; j < sc+SubSize; j++ {
			if i != r || j != c {
				glyph := puz[i][j]
				if glyph != Unknown && glyph != 0 {
					candidates[glyph] = false
				}
			}
		}
	}
	// Return all candidates not eliminated.
	for _, glyph := range Glyphs {
		if candidates[glyph] {
			result = append(result, glyph)
		}
	}
	return
}

// SolveSolo solves a given cell, if it can be determined by simple candidate
// elimination, and then writes to the channel to indicate whether the cell was
// solved.
func (puz *Puzzle) solveSolo(r, c int, ch chan bool) {
	candidates := puz.Candidates(r, c)
	if len(candidates) == 1 {
		puz[r][c] = candidates[0]
		ch <- true
	} else {
		ch <- false
	}
}

// SolveSolos solves all cells in a puzzle which can be found by simple
// candidate elimination.  For each cell which only has one candidate glyph,
// populate the cell with the candidate and repeat until either no unknown
// cells remain, or all remaining unknown cells have multiple candidates.
//
// Return the number of unknown cells remaining.
func (puz *Puzzle) SolveSolos() (remain int) {
	remain = puz.NumUnknowns()
	for {
		if remain == 0 {
			return
		}
		curr := remain
		ch := make(chan bool)
		for i := 0; i < Size; i++ {
			for j := 0; j < Size; j++ {
				if puz[i][j] != Unknown {
					continue
				}
				go puz.solveSolo(i, j, ch)
			}
		}
		for i := 0; i < curr; i++ {
			if <-ch {
				remain--
			}
		}
		if curr == remain {
			// No progress, we've done all we can here ...
			return
		}
	}
	return
}
