package sudoku

// Candidates returns all of the candidate glyphs for a given puzzle cell.
func (puz *Puzzle) Candidates(r, c int) (result []byte) {
	candidates := make(map[byte]bool)
	for _, glyph := range Glyphs {
		candidates[glyph] = true
	}

	// Eliminate glyphs in same row.
	for i := 0; i < Size; i++ {
		if i != c {
			glyph := puz[r][i]
			if Known(glyph) {
				candidates[glyph] = false
			}
		}
	}
	// Eliminate glyphs in same column.
	for i := 0; i < Size; i++ {
		if i != r {
			glyph := puz[i][c]
			if Known(glyph) {
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
				if Known(glyph) {
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

// solveSolo solves a given cell, if it can be determined by simple candidate
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

// glyphInRow returns whether the given glyph is present in the given row.
//
// If 'c' is a valid column index, then that column is excluded from
// consideration.
func (puz *Puzzle) glyphInRow(glyph byte, r, c int) bool {
	for i := 0; i < Size; i++ {
		if i == c {
			continue
		}
		if puz[r][i] == glyph {
			return true
		}
	}
	return false
}

// glyphInColumn returns whether the given glyph is present in the given column.
//
// If 'r' is a valid row index, then that row is excluded from consideration.
func (puz *Puzzle) glyphInColumn(glyph byte, c, r int) bool {
	for i := 0; i < Size; i++ {
		if i == r {
			continue
		}
		if puz[i][c] == glyph {
			return true
		}
	}
	return false
}

// glyphInSubGrid returns whether the given glyph is present in the given
// subgrid index.
//
// If 'r' or 'c' specifies a valid row or column, respectively, in the subgrid,
// then that row or column is excluded from consideration.  These exclusions
// can be disabled by specifying a row or column that is not valid for the
// subgrid (negative numbers will never be valid, so -1 is a good choice here).
func (puz *Puzzle) glyphInSubGrid(glyph byte, subgrid, r, c int) bool {
	sr := (subgrid / 3) * 3
	for i := sr; i < sr+SubSize; i++ {
		if i == r {
			continue
		}
		sc := (subgrid % 3) * 3
		for j := sc; j < sc+SubSize; j++ {
			if j == c {
				continue
			}
			if puz[i][j] == glyph {
				return true
			}
		}
	}
	return false
}

// solveRow finds the location of a glyph within a given row, if all other
// candidate locations for the glyph within the row are invalid.  If the
// location of the glyph is found, it is populated in the puzzle.
//
// The channel receives a boolean to indicate whether the glyph was found.
func (puz *Puzzle) solveRow(glyph byte, r int, ch chan bool) {
	var locs [Size]bool
	var num int
	for i := 0; i < Size; i++ {
		if puz[r][i] == glyph {
			// Glyph is already present in this row, quit.
			ch <- false
			return
		}
		if !Known(puz[r][i]) {
			// Candidate location.  Look for the glyph elsewhere in this column
			// to see whether it can be ruled out.
			if !puz.glyphInColumn(glyph, i, r) {
				locs[i] = true
				num++
			}
		}
	}

	if num > 0 {
		subgrid := (r / SubSize) * SubSize
		sc := 0 // Leftmost column of subgrid
		for i := 0; i < SubSize; i++ {
			if locs[sc] || locs[sc+1] || locs[sc+2] {
				// At least one candidate location in this subgrid.  Search for
				// the target glyph elsewhere in the subgrid to see whether it
				// can be ruled out.
				if puz.glyphInSubGrid(glyph, subgrid, r, -1) {
					// Glyph found, rule out all three locations in the
					// subgrid.
					for j := 0; j < SubSize; j++ {
						if locs[sc+j] {
							locs[sc+j] = false
							num--
						}
					}
					break
				}
			}
			subgrid++
			sc += SubSize
		}
	}

	if num == 1 {
		for i := 0; i < Size; i++ {
			if locs[i] {
				puz[r][i] = glyph
				ch <- true
				return
			}
		}
	}
	ch <- false
}

// SolveSolos solves all cells in a puzzle which can be found by simple
// candidate elimination.  For each cell which only has one candidate glyph, or
// zone with only one candidate location for a glyph, populate the cell with
// the candidate and repeat until either no unknown cells remain, or all
// remaining unknown cells have multiple candidates.
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
