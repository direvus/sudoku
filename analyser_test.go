package sudoku

import "testing"

func TestNumSolutions(t *testing.T) {
	// “Easy” difficulty with a single solution
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
	expect := 1
	solutions := puz.NumSolutions()
	if solutions != expect {
		t.Errorf("incorrect return from NumSolutions(): expected %v, got %v", expect, solutions)
	}

	// “Tricky” difficulty with a single solution
	puz = Puzzle{
		{' ', ' ', '3', ' ', '5', ' ', '2', ' ', ' '},
		{'2', ' ', ' ', '7', ' ', '6', ' ', ' ', '9'},
		{'7', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '4'},
		{' ', '2', ' ', '8', ' ', '1', ' ', '6', ' '},
		{' ', ' ', '9', '6', ' ', '2', '4', ' ', ' '},
		{' ', '4', ' ', '3', ' ', '5', ' ', '2', ' '},
		{'4', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '8'},
		{'3', ' ', ' ', '4', ' ', '8', ' ', ' ', '2'},
		{' ', ' ', '5', ' ', '1', ' ', '3', ' ', ' '}}
	solutions = puz.NumSolutions()
	if solutions != expect {
		t.Errorf("incorrect return from NumSolutions(): expected %v, got %v", expect, solutions)
	}

	// Blank puzzle, multiple solutions
	puz = Puzzle{
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}}
	solutions = puz.NumSolutions()
	if solutions <= 1 {
		t.Errorf("incorrect return from NumSolutions() for blank puzzle: expected multiple, got %v", solutions)
	}

	// Golang challenge 8 sample puzzle, multiple solutions
	puz = Puzzle{
		{'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' '},
		{' ', '5', ' ', ' ', '8', ' ', '1', '2', ' '},
		{'7', ' ', '9', '1', ' ', '3', ' ', '5', '6'},
		{' ', '3', ' ', ' ', '6', '7', ' ', '9', ' '},
		{'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' '},
		{'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7'},
		{' ', '4', ' ', ' ', '7', '8', ' ', '1', ' '},
		{'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' '},
		{' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}}
	solutions = puz.NumSolutions()
	if solutions <= 1 {
		t.Errorf("incorrect return from NumSolutions() for golang-8 puzzle: expected multiple, got %v", solutions)
	}
}
