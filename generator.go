package sudoku

import (
	"time"
	"math/rand"
)

// attemptSolution tries to randomly generate a valid sudoku solution.
//
// It does this by initially populating a small number of cells with random
// values, then attempting to solve the puzzle using a combination of logical
// elimination and brute-force guesswork.
//
// Regardless of whether it succeeds or fails, it writes a pointer to the
// generated Puzzle object to the channel.
//
// This function makes no guarantees as to the correctness or completeness of
// its attempt to generate a solution.  The resulting puzzle may be
// incompletely solved, invalid, or both.  Callers should be prepared to run
// attemptSolution multiple times and discard bad results.
func attemptSolution(ch chan *Puzzle) {
	var puzzle Puzzle
	var glyphs [Size]byte
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	swapper := func(i, j int) {
		glyphs[i], glyphs[j] = glyphs[j], glyphs[i]
	}

	// Populate three randomly selected rows of the puzzle with a shuffled copy
	// of the glyph set.
	copy(glyphs[:], Glyphs[:])
	for i := 0; i < SubSize; i++ {
		// Find a random row that hasn't been populated
		r := 0
		for {
			r = random.Intn(Size)
			if !Known(puzzle[r][0]) {
				break
			}
		}

		// Fill the row with randomly re-ordered glyph values and repeat as
		// necessary if the solution becomes invalid as a result.  It is
		// possible to get into an unresolvable situation here, so put an
		// arbitrary upper limit on the number of attempts.
		for i := 0; i < 100000; i++ {
			random.Shuffle(len(glyphs), swapper)
			copy(puzzle[r][:], glyphs[:])
			if puzzle.Validate() == nil {
				break
			}
		}
	}

	// Attempt to solve the puzzle using simple elimination.  If the puzzle
	// remains unsolved, attempt to solve it using guesswork, starting at a
	// randomly selected unknown cell.
	puzzle.SolveEasy()
	if puzzle.NumUnknowns() == 0 {
		return
	}
	r := random.Intn(Size)
	c := random.Intn(Size)
	found := false
	r, c, found = puzzle.NextUnknown(r, c)
	if !found && (r != 0 || c != 0) {
		r, c, found = puzzle.NextUnknown(0, 0)
	}
	guess := make(chan bool)
	go puzzle.guess(r, c, guess)
	<-guess
	ch <- &puzzle
	return
}

// GenerateSolution returns a randomly generated sudoku solution.
//
// It does this by repeatedly calling attemptSolution(), and returning the
// first puzzle which is both complete and valid.
func GenerateSolution() (puzzle Puzzle) {
	ch := make(chan *Puzzle)
	for {
		go attemptSolution(ch)
		puzzle = *<-ch
		if puzzle.Validate() == nil && puzzle.NumUnknowns() == 0 {
			break
		}
	}
	return
}
