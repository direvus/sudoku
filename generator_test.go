package sudoku

import "testing"

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
