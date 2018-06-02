package sudoku

import (
	"math/rand"
	"time"
)

const MIN_CLUES = 17

// newRand() returns a new random generator initialised with the current time.
func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

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
		random := newRand()
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
	random := newRand()
	r := random.Intn(Size)
	c := random.Intn(Size)
	r, c, _ = puz.FindUnknown(r, c)
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
		puz.Clear()
		go puz.SeedSolution(n, count)
		if <-count == n {
			go puz.AttemptSolution(success)
			if <-success {
				break
			}
		}
	}
	return
}

// MinimalMask returns a minimal clue mask for the given solution.
//
// A minimal clue mask is one from which no clues can be removed without
// causing the puzzle to have multiple solutions.  Each true value in the mask
// indicates the position of a clue in the puzzle, while each false value
// indicates a hidden cell.
func (puz *Puzzle) MinimalMask() (mask Mask) {
	var sol Puzzle
	sol.Merge(*puz)
	knowns := sol.Knowns()
	swapper := func(i, j int) {
		knowns[i], knowns[j] = knowns[j], knowns[i]
	}
	count := len(knowns)
	random := newRand()
	for count >= MIN_CLUES {
		random.Shuffle(count, swapper)
		found := false
		for i, cell := range knowns {
			var attempt Puzzle
			attempt.Merge(sol)
			attempt.SetCell(cell, Unknown)
			n := attempt.NumSolutions()
			if n == 1 {
				// So far so good.  Drop the cell from the solution and start
				// the next pass.
				sol.SetCell(cell, Unknown)
				swapper(i, count-1)
				knowns = knowns[:count-1]
				count--
				found = true
				break
			}
		}
		if !found {
			// None of the known cells could be safely removed.  Exit out.
			break
		}
	}
	mask = sol.GetMask()
	return
}
