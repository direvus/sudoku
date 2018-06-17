package main

import (
	"github.com/direvus/sudoku"
	"os"
)

func main() {
	solution := sudoku.GenerateSolution()
	mask := solution.MinimalMask()
	puzzle := solution.ApplyMask(&mask)

	os.Stdout.WriteString(puzzle.String())
	os.Exit(0)
}
