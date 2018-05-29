package sudoku

import "testing"

func TestSeedSolution(t *testing.T) {
	var puz Puzzle
	ch := make(chan int)
	n := 1
	go puz.SeedSolution(n, ch)
	result := <-ch
	if result != n {
		t.Errorf("incorrect count from SeedSolution(): expected %v cells populated, got %v", n, result)
	}

	err := puz.Validate()
	if err != nil {
		t.Errorf("invalid result from SeedSolution(): %v\n\n%v", err, puz.String())
	}

	puz.Clear()
	n = 27
	go puz.SeedSolution(n, ch)
	result = <-ch
	if result > n || result < 1 {
		t.Errorf("unexpected count from SeedSolution(): expected at least 1 and up to %v cells populated, got %v", n, result)
	}

	// Attempt to seed a complete puzzle
	puz = Puzzle{
		{'2', '9', '7', '6', '3', '5', '8', '1', '4'},
		{'6', '5', '1', '4', '2', '8', '7', '9', '3'},
		{'4', '8', '3', '1', '9', '7', '5', '2', '6'},
		{'7', '4', '8', '5', '1', '9', '6', '3', '2'},
		{'9', '6', '5', '2', '7', '3', '1', '4', '8'},
		{'1', '3', '2', '8', '4', '6', '9', '7', '5'},
		{'5', '1', '9', '3', '6', '2', '4', '8', '7'},
		{'3', '7', '6', '9', '8', '4', '2', '5', '1'},
		{'8', '2', '4', '7', '5', '1', '3', '6', '9'}}
	go puz.SeedSolution(n, ch)
	result = <-ch
	if result != 0 {
		t.Errorf("unexpected count from SeedSolution(): expected 0, got %v", result)
	}

	// Attempt to seed a puzzle containing a contradiction
	puz = Puzzle{
		{'2', '9', '7', '6', '3', '5', '8', '1', '4'},
		{'6', '5', '1', '4', '2', '8', '7', '9', '3'},
		{'4', '8', '3', '1', '9', '7', '5', '2', '6'},
		{'7', '4', '8', '5', '1', '9', '6', '3', '2'},
		{'9', '6', '5', '2', ' ', '3', '1', '7', '8'},
		{'1', '3', '2', '8', '4', '6', '9', '4', '5'},
		{'5', '1', '9', '3', '6', '2', '4', '8', '7'},
		{'3', '7', '6', '9', '8', '4', '2', '5', '1'},
		{'8', '2', '4', '7', '5', '1', '3', '6', '9'}}
	go puz.SeedSolution(n, ch)
	result = <-ch
	if result != 0 {
		t.Errorf("unexpected count from SeedSolution(): expected 0, got %v", result)
	}
	if Known(puz[4][4]) {
		t.Errorf("unexpected value in R5C5 from SeedSolution(): expected unknown, got %q", puz[4][4])
	}
}

func TestAttemptSolution(t *testing.T) {
	ch := make(chan bool)
	// “Easy” difficulty
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
	go puz.AttemptSolution(ch)
	if !<-ch {
		t.Errorf("AttemptSolution indicated failure:\n\n%v", puz.String())
	}
	if !puz.Equal(expect) {
		t.Errorf("unexpected result from AttemptSolution: expected:\n%v\n\ngot:\n%v", expect.String(), puz.String())
	}

	// “Tricky” difficulty
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
	go puz.AttemptSolution(ch)
	if !<-ch {
		t.Errorf("AttemptSolution indicated failure:\n\n%v", puz.String())
	}
	remain := puz.NumUnknowns()
	if remain != 0 {
		t.Errorf("too many unknowns after AttemptSolution: expected %v unknowns remaining, got %v", 0, remain)
	}
	err := puz.Validate()
	if err != nil {
		t.Errorf("invalid puzzle from AttemptSolution: %v\n\n%v", err, puz.String())
	}

	// “Extreme” difficulty
	puz = Puzzle{
		{' ', ' ', '8', ' ', ' ', '6', '2', '5', ' '},
		{' ', ' ', ' ', ' ', '7', ' ', ' ', '3', ' '},
		{' ', ' ', ' ', ' ', '1', '2', '9', '8', ' '},
		{' ', ' ', '5', ' ', ' ', '3', ' ', ' ', ' '},
		{' ', '2', ' ', '7', ' ', '1', ' ', '6', ' '},
		{' ', ' ', ' ', '8', ' ', ' ', '1', ' ', ' '},
		{' ', '3', '6', '2', '8', ' ', ' ', ' ', ' '},
		{' ', '7', ' ', ' ', '9', ' ', ' ', ' ', ' '},
		{' ', '8', '2', '1', ' ', ' ', '4', ' ', ' '}}
	go puz.AttemptSolution(ch)
	if !<-ch {
		t.Errorf("AttemptSolution indicated failure:\n\n%v", puz.String())
	}
	remain = puz.NumUnknowns()
	if remain != 0 {
		t.Errorf("too many unknowns after AttemptSolution: expected %v unknowns remaining, got %v", 0, remain)
	}
	err = puz.Validate()
	if err != nil {
		t.Errorf("invalid puzzle from AttemptSolution: %v\n\n%v", err, puz.String())
	}
}

func TestGenerateSolution(t *testing.T) {
	puz := GenerateSolution()
	unknowns := puz.NumUnknowns()
	if unknowns != 0 {
		t.Errorf("incorrect result from GenerateSolution(): expected zero unknowns, got %v:\n\n%v", unknowns, puz.String())
	}

	err := puz.Validate()
	if err != nil {
		t.Errorf("invalid result from GenerateSolution(): %v\n\n%v", err, puz.String())
	}
}
