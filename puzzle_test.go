package sudoku

import (
	"bytes"
	"testing"
)

func TestKnown(t *testing.T) {
	tests := []byte{0, ' ', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	expect := []bool{false, false, true, true, true, true, true, true, true, true, true}
	for i, v := range tests {
		result := Known(v)
		if result != expect[i] {
			t.Errorf("invalid result from Known for %v: expected %v, got %v", v, expect[i], result)
		}
	}
}

func TestPuzzleRead(t *testing.T) {
	var puz Puzzle
	var err error

	// Valid puzzle input
	err = puz.Read([]byte(
		"1 _ 3 _ _ 6 _ 8 _\n" +
			"_ 5 _ _ 8 _ 1 2 _\n" +
			"7 _ 9 1 _ 3 _ 5 6\n" +
			"_ 3 _ _ 6 7 _ 9 _\n" +
			"5 _ 7 8 _ _ _ 3 _\n" +
			"8 _ 1 _ 3 _ 5 _ 7\n" +
			"_ 4 _ _ 7 8 _ 1 _\n" +
			"6 _ 8 _ _ 2 _ 4 _\n" +
			"_ 1 2 _ 4 5 _ 7 8\n"))
	if err != nil {
		t.Errorf("Read failed on valid input: %s", err)
	}

	// Invalid puzzle input (prose)
	err = puz.Read([]byte("This isn't sudoku!"))
	if err == nil {
		t.Errorf("no error for malformed input")
	}

	// Invalid puzzle input (too few lines)
	err = puz.Read([]byte(
		"1 _ 3 _ _ 6 _ 8 _\n" +
			"_ 5 _ _ 8 _ 1 2 _\n"))
	if err == nil {
		t.Errorf("no error for insufficient lines")
	}

	// Invalid puzzle input (junk characters around puzzle)
	err = puz.Read([]byte(
		"a1 _ 3 _ _ 6 _ 8 _\n" +
			"_ 5 _ _ 8 _ 1 2 _\n" +
			"7 _ 9 1 _ 3 _ 5 6\n" +
			"_ 3 _ _ 6 7 _ 9 _\n" +
			"5 _ 7 8 _ _ _ 3 _\n" +
			"8 _ 1 _ 3 _ 5 _ 7\n" +
			"_ 4 _ _ 7 8 _ 1 _\n" +
			"6 _ 8 _ _ 2 _ 4 _\n" +
			"_ 1 2 _ 4 5 _ 7 8z\n"))
	if err == nil {
		t.Errorf("no error for junk characters around input")
	}

	// Invalid puzzle input (invalid glyph in puzzle)
	err = puz.Read([]byte(
		"1 _ 3 _ _ 6 _ 8 _\n" +
			"_ 5 _ _ 8 _ 1 2 _\n" +
			"7 _ 9 1 _ 3 _ 5 6\n" +
			"_ 3 _ _ 6 7 _ 9 _\n" +
			"5 _ 7 8 _ _ _ 3 _\n" +
			"8 _ 1 _ A _ 5 _ 7\n" +
			"_ 4 _ _ 7 8 _ 1 _\n" +
			"6 _ 8 _ _ 2 _ 4 _\n" +
			"_ 1 2 _ 4 5 _ 7 8\n"))
	if err == nil {
		t.Errorf("no error for invalid glyph in row 6")
	}

	// Invalid puzzle input (too few glyphs)
	err = puz.Read([]byte(
		"1 _ 3 _ _ 6 _ 8 _\n" +
			"_ 5 _ _ 8 _ 1 2 _\n" +
			"7 _ 9 1 _ 3 _ 5 6\n" +
			"_ 3 _ _ 6 7 _ 9 _\n" +
			"5 _ 7 8 _ _ _ 3 _\n" +
			"8 _ 1 _ 3 _ 5 _\n" +
			"_ 4 _ _ 7 8 _ 1 _\n" +
			"6 _ 8 _ _ 2 _ 4 _\n" +
			"_ 1 2 _ 4 5 _ 7 8\n"))
	if err == nil {
		t.Errorf("no error for insufficient glyphs in row 6")
	}
}

func TestPuzzleRow(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	r := 0
	expect := []byte{'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' '}
	row := puz.Row(r)
	if !bytes.Equal(row, expect) {
		t.Errorf("incorrect value for row %v, expected %v, got %v", r+1, expect, row)
	}

	r = 1
	expect = []byte{' ', '5', ' ', ' ', '8', ' ', '1', '2', ' '}
	row = puz.Row(r)
	if !bytes.Equal(row, expect) {
		t.Errorf("incorrect value for row %v, expected %v, got %v", r+1, expect, row)
	}

	r = 8
	expect = []byte{' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	row = puz.Row(r)
	if !bytes.Equal(row, expect) {
		t.Errorf("incorrect value for row %v, expected %v, got %v", r+1, expect, row)
	}
}

func TestPuzzleColumn(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	c := 0
	expect := []byte{'1', ' ', '7', ' ', '5', '8', ' ', '6', ' '}
	column := puz.Column(c)
	if !bytes.Equal(column, expect) {
		t.Errorf("incorrect value for column %v, expected %v, got %v", c+1, expect, column)
	}

	c = 1
	expect = []byte{' ', '5', ' ', '3', ' ', ' ', '4', ' ', '1'}
	column = puz.Column(c)
	if !bytes.Equal(column, expect) {
		t.Errorf("incorrect value for column %v, expected %v, got %v", c+1, expect, column)
	}

	c = 8
	expect = []byte{' ', ' ', '6', ' ', ' ', '7', ' ', ' ', '8'}
	column = puz.Column(c)
	if !bytes.Equal(column, expect) {
		t.Errorf("incorrect value for column %v, expected %v, got %v", c+1, expect, column)
	}
}

func TestCellSubGrid(t *testing.T) {
	expect := [Size][Size]int{
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8}}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			result := CellSubGrid(i, j)
			if result != expect[i][j] {
				t.Errorf("invalid return from CellSubGrid for R%vC%v, expected %v, got %v", i, j, expect[i][j], result)
			}
		}
	}
}

func TestPuzzleSubGrid(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	g := 0
	expect := []byte{'1', ' ', '3', ' ', '5', ' ', '7', ' ', '9'}
	subgrid := puz.SubGrid(g)
	if !bytes.Equal(subgrid, expect) {
		t.Errorf("incorrect value for subgrid %v, expected %q, got %q", g+1, expect, subgrid)
	}

	g = 1
	expect = []byte{' ', ' ', '6', ' ', '8', ' ', '1', ' ', '3'}
	subgrid = puz.SubGrid(g)
	if !bytes.Equal(subgrid, expect) {
		t.Errorf("incorrect value for subgrid %v, expected %q, got %q", g+1, expect, subgrid)
	}

	g = 8
	expect = []byte{' ', '1', ' ', ' ', '4', ' ', ' ', '7', '8'}
	subgrid = puz.SubGrid(g)
	if !bytes.Equal(subgrid, expect) {
		t.Errorf("incorrect value for subgrid %v, expected %q, got %q", g+1, expect, subgrid)
	}
}

func TestPuzzleEqual(t *testing.T) {
	a := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	b := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '9'}
	c := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	if a.Equal(b) {
		t.Errorf("incorrect result from Equal: %v is not equal to %v", a, b)
	}
	if b.Equal(c) {
		t.Errorf("incorrect result from Equal: %v is not equal to %v", a, b)
	}
	if !a.Equal(c) {
		t.Errorf("incorrect result from Equal: %v is equal to %v", a, c)
	}
}

func TestPuzzleNumUnknowns(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := 40
	result := puz.NumUnknowns()
	if result != expect {
		t.Errorf("incorrect return from NumUnknown: expected %v, got %v", expect, result)
	}

	puz = Puzzle{
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	expect = 81
	result = puz.NumUnknowns()
	if result != expect {
		t.Errorf("incorrect return from NumUnknown: expected %v, got %v", expect, result)
	}

	puz = Puzzle{
		'1', '2', '3', '4', '5', '6', '7', '8', '9',
		'4', '5', '6', '7', '8', '9', '1', '2', '3',
		'7', '8', '9', '1', '2', '3', '4', '5', '6',
		'2', '3', '4', '5', '6', '7', '8', '9', '1',
		'5', '6', '7', '8', '9', '1', '2', '3', '4',
		'8', '9', '1', '2', '3', '4', '5', '6', '7',
		'3', '4', '5', '6', '7', '8', '9', '1', '2',
		'6', '7', '8', '9', '1', '2', '3', '4', '5',
		'9', '1', '2', '3', '4', '5', '6', '7', '8'}
	expect = 0
	result = puz.NumUnknowns()
	if result != expect {
		t.Errorf("incorrect return from NumUnknown: expected %v, got %v", expect, result)
	}
}

func TestPuzzleNextUnknown(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	r, c, found := puz.NextUnknown(0, 0)
	if r != 0 || c != 1 || !found {
		t.Errorf("incorrect return from NextUnknown: expected R1C2, got R%vC%v", r+1, c+1)
	}

	r, c, found = puz.NextUnknown(5, 8)
	if r != 6 || c != 0 || !found {
		t.Errorf("incorrect return from NextUnknown: expected R7C1, got R%vC%v", r+1, c+1)
	}

	r, c, found = puz.NextUnknown(8, 7)
	if found {
		t.Errorf("incorrect return from NextUnknown: expected none found, got R%vC%v", r+1, c+1)
	}

	puz = Puzzle{
		'1', '2', '3', '4', '5', '6', '7', '8', '9',
		'4', '5', '6', '7', '8', '9', '1', '2', '3',
		'7', '8', '9', '1', '2', '3', '4', '5', '6',
		'2', '3', '4', '5', '6', '7', '8', '9', '1',
		'5', '6', '7', '8', '9', '1', '2', '3', '4',
		'8', '9', '1', '2', '3', '4', '5', '6', '7',
		'3', '4', '5', '6', '7', '8', '9', '1', '2',
		'6', '7', '8', '9', '1', '2', '3', '4', '5',
		'9', '1', '2', '3', '4', '5', '6', '7', '8'}
	r, c, found = puz.NextUnknown(0, 0)
	if found {
		t.Errorf("incorrect return from NextUnknown: expected none, got R%vC%v", r+1, c+1)
	}
}

func TestPuzzleFindUnknown(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	r, c, found := puz.FindUnknown(0, 0)
	if r != 0 || c != 1 || !found {
		t.Errorf("incorrect return from FindUnknown: expected R1C2, got R%vC%v", r+1, c+1)
	}

	r, c, found = puz.FindUnknown(5, 8)
	if r != 6 || c != 0 || !found {
		t.Errorf("incorrect return from FindUnknown: expected R7C1, got R%vC%v", r+1, c+1)
	}

	r, c, found = puz.FindUnknown(8, 7)
	if r != 0 || c != 1 || !found {
		t.Errorf("incorrect return from FindUnknown: expected R1C2, got R%vC%v", r+1, c+1)
	}

	puz = Puzzle{
		'1', '2', '3', '4', '5', '6', '7', '8', '9',
		'4', '5', '6', '7', '8', '9', '1', '2', '3',
		'7', '8', '9', '1', '2', '3', '4', '5', '6',
		'2', '3', '4', '5', '6', '7', '8', '9', '1',
		'5', '6', '7', '8', '9', '1', '2', '3', '4',
		'8', '9', '1', '2', '3', '4', '5', '6', '7',
		'3', '4', '5', '6', '7', '8', '9', '1', '2',
		'6', '7', '8', '9', '1', '2', '3', '4', '5',
		'9', '1', '2', '3', '4', '5', '6', '7', '8'}
	r, c, found = puz.FindUnknown(4, 4)
	if found {
		t.Errorf("incorrect return from FindUnknown: expected none, got R%vC%v", r+1, c+1)
	}
}

func TestPuzzleMerge(t *testing.T) {
	a := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	b := Puzzle{
		  0,  '2',	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0,
		  0,  '8',	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0, '6',   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0,
		  0,	0,	 0,   0,   0,	0,	 0,   0,   0}
	expect := Puzzle{
		'1', '2', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', '8', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', '6', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	a.Merge(b)
	if !a.Equal(expect) {
		t.Errorf("incorrect result from merge: expected %v, got %v", expect, a)
	}
}

func TestPuzzleApplyMask(t *testing.T) {
	p := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	var m Mask
	// No cells masked
	result := p.ApplyMask(&m)
	expect := Puzzle{
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	if !result.Equal(expect) {
		t.Errorf("incorrect result from ApplyMask: expected\n%v\n\ngot\n%v", expect.String(), result.String())
	}

	// R7C5 only masked
	m[58] = true
	expect = Puzzle{
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', '7', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

	result = p.ApplyMask(&m)
	if !result.Equal(expect) {
		t.Errorf("incorrect result from ApplyMask: expected\n%v\n\ngot\n%v", expect.String(), result.String())
	}

	// Everything masked
	m = Mask{
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true}
	expect = Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	result = p.ApplyMask(&m)
	if !result.Equal(expect) {
		t.Errorf("incorrect result from ApplyMask: expected\n%v\n\ngot\n%v", expect.String(), result.String())
	}
}

func TestFindDuplicate(t *testing.T) {
	tests := [][]byte{
		{'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' '},
		{'1', ' ', '3', ' ', ' ', '6', ' ', '8', '1'},
		{' ', '5', ' ', ' ', '8', ' ', '1', '2', ' '},
		{' ', '5', ' ', ' ', '8', ' ', '1', '2', '2'},
		{'8', '5', ' ', ' ', '8', ' ', '1', '2', '2'}}
	expect := []byte{
		0,
		'1',
		0,
		'2',
		'8'}
	for i, test := range tests {
		result := findDuplicate(test)
		if result != expect[i] {
			t.Errorf("incorrect findDuplicate result: expected %q, got %q", expect[i], result)
		}
	}
}

func TestPuzzleValidate(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	var err error
	var orig byte
	err = puz.Validate()
	if err != nil {
		t.Errorf("error for valid puzzle: %v", err)
	}

	// Duplicate value in row
	orig, puz[32] = puz[32], '9'
	err = puz.Validate()
	if err == nil {
		t.Errorf("no error for duplicate glyph in row 4")
	}
	puz[32] = orig

	// Duplicate value in column
	orig, puz[0] = puz[0], '7'
	err = puz.Validate()
	if err == nil {
		t.Errorf("no error for duplicate glyph in column 1")
	}
	puz[0] = orig

	puz2 := Puzzle{
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		'3', '9', '7', '5', '6', '2', '4', '8', '1',
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		'9', '8', '3', '6', '5', '7', '4', '1', '2',
		  0,   0,   0,   0,   0,   0,   0,   0,   0}
	err = puz2.Validate()
	if err == nil {
		t.Errorf("no error for duplicate glyph in column 7")
	}

	// Duplicate value in subgrid
	orig, puz[71] = puz[71], '1'
	err = puz.Validate()
	if err == nil {
		t.Errorf("no error for duplicate glyph in subgrid 9")
	}
	puz[71] = orig
}

func TestPuzzleString(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := (
		"1 _ 3 _ _ 6 _ 8 _\n" +
		"_ 5 _ _ 8 _ 1 2 _\n" +
		"7 _ 9 1 _ 3 _ 5 6\n" +
		"_ 3 _ _ 6 7 _ 9 _\n" +
		"5 _ 7 8 _ _ _ 3 _\n" +
		"8 _ 1 _ 3 _ 5 _ 7\n" +
		"_ 4 _ _ 7 8 _ 1 _\n" +
		"6 _ 8 _ _ 2 _ 4 _\n" +
		"_ 1 2 _ 4 5 _ 7 8\n")
	result := puz.String()
	if expect != result {
		t.Errorf("invalid string output, expected\n%v\n\ngot\n%v", expect, result)
	}
}

func TestPuzzleClear(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := Puzzle{
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0,
		  0,   0,   0,   0,   0,   0,   0,   0,   0}
	puz.Clear()
	if !puz.Equal(expect) {
		t.Errorf("invalid result from Clear(), expected all zero bytes, got:\n\n%v", puz.String())
	}
}

func TestPuzzleUnknowns(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := []CellRef{
		{0, 1}, {0, 3}, {0, 4}, {0, 6}, {0, 8},
		{1, 0}, {1, 2}, {1, 3}, {1, 5}, {1, 8},
		{2, 1}, {2, 4}, {2, 6},
		{3, 0}, {3, 2}, {3, 3}, {3, 6}, {3, 8},
		{4, 1}, {4, 4}, {4, 5}, {4, 6}, {4, 8},
		{5, 1}, {5, 3}, {5, 5}, {5, 7},
		{6, 0}, {6, 2}, {6, 3}, {6, 6}, {6, 8},
		{7, 1}, {7, 3}, {7, 4}, {7, 6}, {7, 8},
		{8, 0}, {8, 3}, {8, 6}}
	result := puz.Unknowns()
	if len(result) != len(expect) {
		t.Errorf("incorrect return from Unknowns: expected %v cells, got %v", len(expect), len(result))
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Errorf("incorrect return from Unknowns: expected %v at index %v, got %v", expect[i], i, result[i])
		}
	}
}

func TestPuzzleKnowns(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := []CellRef{
		{0, 0}, {0, 2}, {0, 5}, {0, 7},
		{1, 1}, {1, 4}, {1, 6}, {1, 7},
		{2, 0}, {2, 2}, {2, 3}, {2, 5}, {2, 7}, {2, 8},
		{3, 1}, {3, 4}, {3, 5}, {3, 7},
		{4, 0}, {4, 2}, {4, 3}, {4, 7},
		{5, 0}, {5, 2}, {5, 4}, {5, 6}, {5, 8},
		{6, 1}, {6, 4}, {6, 5}, {6, 7},
		{7, 0}, {7, 2}, {7, 5}, {7, 7},
		{8, 1}, {8, 2}, {8, 4}, {8, 5}, {8, 7}, {8, 8}}
	result := puz.Knowns()
	if len(result) != len(expect) {
		t.Errorf("incorrect return from Knowns: expected %v cells, got %v", len(expect), len(result))
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Errorf("incorrect return from Knowns: expected %v at index %v, got %v", expect[i], i, result[i])
		}
	}
}

func TestPuzzleGetCell(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	cell := CellRef{5, 6}
	v := puz.GetCell(cell)
	expect := byte('5')
	if v != expect {
		t.Errorf("incorrect result from GetCell: expected %v, got %v", expect, v)
	}
}

func TestPuzzleSetCell(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	cell := CellRef{5, 7}
	puz.SetCell(cell, '6')
	expect := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', '6', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	if !puz.Equal(expect) {
		t.Errorf("incorrect result from SetCell: expected\n%v\n\ngot\n%v", expect.String(), puz.String())
	}
}

func TestPuzzleGetMask(t *testing.T) {
	puz := Puzzle{
		'1', ' ', '3', ' ', ' ', '6', ' ', '8', ' ',
		' ', '5', ' ', ' ', '8', ' ', '1', '2', ' ',
		'7', ' ', '9', '1', ' ', '3', ' ', '5', '6',
		' ', '3', ' ', ' ', '6', '7', ' ', '9', ' ',
		'5', ' ', '7', '8', ' ', ' ', ' ', '3', ' ',
		'8', ' ', '1', ' ', '3', ' ', '5', ' ', '7',
		' ', '4', ' ', ' ', '7', '8', ' ', '1', ' ',
		'6', ' ', '8', ' ', ' ', '2', ' ', '4', ' ',
		' ', '1', '2', ' ', '4', '5', ' ', '7', '8'}
	expect := Mask{
		 true, false,  true, false, false,  true, false,  true, false,
		false,  true, false, false,  true, false,  true,  true, false,
		 true, false,  true,  true, false,  true, false,  true,  true,
		false,  true, false, false,  true,  true, false,  true, false,
		 true, false,  true,  true, false, false, false,  true, false,
		 true, false,  true, false,  true, false,  true, false,  true,
		false,  true, false, false,  true,  true, false,  true, false,
		 true, false,  true, false, false,  true, false,  true, false,
		false,  true,  true, false,  true,  true, false,  true,  true}
	mask := puz.GetMask()
	if !mask.Equal(expect) {
		t.Errorf("incorrect result from GetMask: expected\n%v\n\ngot\n%v", expect.String(), mask.String())
	}
}
