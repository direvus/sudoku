package sudoku

import "testing"

func TestCellRefString(t *testing.T) {
	refs := []struct{
		row int
		col int
		expect string
	}{
		{0, 0, "R1C1"},
	}
	for i := 0; i < len(refs); i++ {
		ref := CellRef{refs[i].row, refs[i].col}
		result := ref.String()
		if refs[i].expect != result {
			t.Errorf("invalid string output, expected\n%v\n\ngot\n%v", refs[i].expect, result)
		}
	}
}
