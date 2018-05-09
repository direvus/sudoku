package sudoku

import (
	"bytes"
	"testing"
)

func TestCandidates(t *testing.T) {
	puz := Puzzle{
		{'2', ' ', ' ', '6', '3', ' ', ' ', '1', ' '},
		{' ', '5', '1', ' ', '2', ' ', '7', '9', '3'},
		{'4', ' ', '3', '1', '9', '7', '5', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '9', ' ', '3', '2'},
		{' ', '6', '5', ' ', '7', ' ', '1', '4', ' '},
		{'1', '3', ' ', '8', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', '9', '3', '6', '2', '4', ' ', '7'},
		{'3', '7', '6', ' ', '8', ' ', '2', '5', ' '},
		{' ', '2', ' ', ' ', '5', '1', ' ', ' ', '9'}}
	expect := []byte{'1'}
	result := puz.Candidates(7, 8)
	if !bytes.Equal(result, expect) {
		t.Errorf("incorrect candidates, expected %q, got %q", expect, result)
	}

	expect = []byte{'4', '5', '6'}
	result = puz.Candidates(5, 5)
	if !bytes.Equal(result, expect) {
		t.Errorf("incorrect candidates, expected %q, got %q", expect, result)
	}
}

func TestSolveSolos(t *testing.T) {
	puz := Puzzle{
		{'2', ' ', ' ', '6', '3', ' ', ' ', '1', ' '},
		{' ', '5', '1', ' ', '2', ' ', '7', '9', '3'},
		{'4', ' ', '3', '1', '9', '7', '5', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '9', ' ', '3', '2'},
		{' ', '6', '5', ' ', '7', ' ', '1', '4', ' '},
		{'1', '3', ' ', '8', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', '9', '3', '6', '2', '4', ' ', '7'},
		{'3', '7', '6', ' ', '8', ' ', '2', '5', ' '},
		{' ', '2', ' ', ' ', '5', '1', ' ', ' ', '9'}}
	expect := Puzzle{
		{'2', '9', '7', '6', '3', '5', '8', '1', '4'},
		{'6', '5', '1', '4', '2', '8', '7', '9', '3'},
		{'4', '8', '3', '1', '9', '7', '5', '2', '6'},
		{'7', '4', '8', '5', '1', '9', '6', '3', '2'},
		{'9', '6', '5', '2', '7', '3', '1', '4', '8'},
		{'1', '3', '2', '8', '4', '6', '9', '7', '5'},
		{'5', '1', '9', '3', '6', '2', '4', '8', '7'},
		{'3', '7', '6', '9', '8', '4', '2', '5', '1'},
		{'8', '2', '4', '7', '5', '1', '3', '6', '9'}}
	remain := puz.SolveSolos()
	if remain != 0 {
		t.Errorf("incorrect return from SolveSolos: expected %v unknowns remaining, got %v", 0, remain)
	}
	if !puz.Equal(expect) {
		t.Errorf("puzzle not solved: expected %v, got %v", expect, puz)
	}
}

func TestSolveRow(t *testing.T) {
	puz := Puzzle{
		{' ', ' ', '3', ' ', '5', ' ', '2', ' ', ' '},
		{'2', ' ', ' ', '7', ' ', '6', ' ', ' ', '9'},
		{'7', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '4'},
		{' ', '2', ' ', '8', ' ', '1', ' ', '6', ' '},
		{' ', ' ', '9', '6', ' ', '2', '4', ' ', ' '},
		{' ', '4', ' ', '3', ' ', '5', ' ', '2', ' '},
		{'4', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '8'},
		{'3', ' ', ' ', '4', ' ', '8', ' ', ' ', '2'},
		{' ', ' ', '5', ' ', '1', ' ', '3', ' ', ' '}}
	ch := make(chan bool)
	go puz.solveRow('2', 8, ch)
	if !<-ch {
		t.Errorf("failed to solve for 2 in row 8")
	}
	if puz[8][3] != '2' {
		t.Errorf("failed to populate solution for 2 in row 8, found %q", puz[8][3])
	}

	go puz.solveRow('1', 0, ch)
	if <-ch {
		t.Errorf("unexpectedly solved for 1 in row 0, multiple candidate locations")
	}
}

func TestSolveColumn(t *testing.T) {
	puz := Puzzle{
		{' ', ' ', '3', ' ', '5', ' ', '2', ' ', ' '},
		{'2', ' ', ' ', '7', ' ', '6', ' ', ' ', '9'},
		{'7', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '4'},
		{' ', '2', ' ', '8', ' ', '1', ' ', '6', ' '},
		{' ', ' ', '9', '6', ' ', '2', '4', ' ', ' '},
		{' ', '4', ' ', '3', ' ', '5', ' ', '2', ' '},
		{'4', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '8'},
		{'3', ' ', ' ', '4', ' ', '8', ' ', ' ', '2'},
		{' ', ' ', '5', ' ', '1', ' ', '3', ' ', ' '}}
	ch := make(chan bool)
	go puz.solveColumn('5', 3, ch)
	if !<-ch {
		t.Errorf("failed to solve for 5 in column 3")
	}
	if puz[6][3] != '5' {
		t.Errorf("failed to populate solution for 5 in column 3, found %q", puz[6][3])
	}

	go puz.solveColumn('1', 0, ch)
	if <-ch {
		t.Errorf("unexpectedly solved for 1 in column 0, multiple candidate locations")
	}
}

func BenchmarkSolveSolos(b *testing.B) {
	puz := Puzzle{
		{'2', ' ', ' ', '6', '3', ' ', ' ', '1', ' '},
		{' ', '5', '1', ' ', '2', ' ', '7', '9', '3'},
		{'4', ' ', '3', '1', '9', '7', '5', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '9', ' ', '3', '2'},
		{' ', '6', '5', ' ', '7', ' ', '1', '4', ' '},
		{'1', '3', ' ', '8', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', '9', '3', '6', '2', '4', ' ', '7'},
		{'3', '7', '6', ' ', '8', ' ', '2', '5', ' '},
		{' ', '2', ' ', ' ', '5', '1', ' ', ' ', '9'}}
	for i := 0; i < b.N; i++ {
		var test Puzzle
		test.Merge(puz)
		test.SolveSolos()
	}
}
