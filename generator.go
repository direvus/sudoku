package sudoku

import (
	"math/rand"
	"time"
)

// SeedSolution populates random glyphs into a sudoku puzzle.
//
// It does this by randomly selecting an unknown cell, and then trying random
// glyphs in that cell.  The first glyph tried that does not cause a
// contradiction is selected for the cell.
//
// The function continues in this manner until it has successfully populated
// 'n' cells, or it detects an unresolvable cell.  In either case, it writes
// the number of cells populated to the channel.
func (puz *Puzzle) SeedSolution(n int, ch chan int) {
	var glyphs [Size]byte
	swapper := func(i, j int) {
		glyphs[i], glyphs[j] = glyphs[j], glyphs[i]
	}
	copy(glyphs[:], Glyphs[:])
	i := 0
	for ; i < n; i++ {
		if puz.NumUnknowns() == 0 {
			ch <- i
			return
		}
		// Select an unknown cell at random
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		var r, c int
		for {
			r := random.Intn(Size)
			c := random.Intn(Size)
			if !Known(puz[r][c]) {
				break
			}
		}
		// Shuffle the glyphs array and try each glyph in turn
		random.Shuffle(len(glyphs), swapper)
		for j := 0; j < len(glyphs); j++ {
			puz[r][c] = glyphs[j]
			if puz.Validate() == nil {
				break
			}
			puz[r][c] = Unknown
		}
		if !Known(puz[r][c]) {
			break
		}
	}
	ch <- i
	return
}

// AttemptSolution tries to randomly generate a valid sudoku solution.
//
// It does this by attempting to solve the puzzle using a combination of
// logical elimination and brute-force guesswork starting from a randomly
// selected unknown cell.
//
// The channel receives true if a solution was found, false otherwise.
func (puz *Puzzle) AttemptSolution(ch chan bool) {
	// Try simple logical elimination.
	puz.SolveEasy()
	if puz.NumUnknowns() == 0 {
		ch <- true
		return
	}
	// Select a random cell to start guessing from.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := random.Intn(Size)
	c := random.Intn(Size)
	found := false
	r, c, found = puz.NextUnknown(r, c)
	if !found && (r != 0 || c != 0) {
		r, c, found = puz.NextUnknown(0, 0)
	}
	guess := make(chan bool)
	go puz.guess(r, c, guess)
	ch <- <-guess
}

// GenerateSolution returns a randomly generated sudoku solution.
//
// It does this by repeatedly seeding an empty puzzle with a set of randomly
// chosen and located glyphs, and then calling AttemptSolution() on the result.
// It returns the first such puzzle which is both complete and valid.
func GenerateSolution() (puz Puzzle) {
	success := make(chan bool)
	count := make(chan int)
	n := 27
	for {
		go puz.SeedSolution(n, count)
		if <-count == n {
			go puz.AttemptSolution(success)
			if <-success {
				break
			}
		}
		puz.Clear()
	}
	return
}
