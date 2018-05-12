package main

import (
	"bytes"
	"os"
	"github.com/direvus/sudoku"
)

func main() {
	var buf bytes.Buffer
	var puzzle sudoku.Puzzle

	buf.ReadFrom(os.Stdin)
	err := puzzle.Read(buf.Bytes())
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	puzzle.Solve()
	os.Stdout.WriteString(puzzle.String())
	os.Exit(0)
}
