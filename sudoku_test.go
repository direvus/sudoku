package sudoku

import "testing"

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

func TestPuzzleValidate(t *testing.T) {
    puz := Puzzle{
        {'1',' ','3',' ',' ','6',' ','8',' '},
        {' ','5',' ',' ','8',' ','1','2',' '},
        {'7',' ','9','1',' ','3',' ','5','6'},
        {' ','3',' ',' ','6','7',' ','9',' '},
        {'5',' ','7','8',' ',' ',' ','3',' '},
        {'8',' ','1',' ','3',' ','5',' ','7'},
        {' ','4',' ',' ','7','8',' ','1',' '},
        {'6',' ','8',' ',' ','2',' ','4',' '},
        {' ','1','2',' ','4','5',' ','7','8'}}
    var err error
    var orig byte
    err = puz.Validate()
    if err != nil {
        t.Errorf("error for valid puzzle: %v", err)
    }

    // Duplicate value in row
    orig, puz[3][5] = puz[3][5], '9'
    err = puz.Validate()
    if err == nil {
        t.Errorf("no error for duplicate glyph in row 4")
    }
    puz[3][5] = orig

    // Duplicate value in column
    orig, puz[0][0] = puz[0][0], '7'
    err = puz.Validate()
    if err == nil {
        t.Errorf("no error for duplicate glyph in column 1")
    }
    puz[0][0] = orig

    // Duplicate value in subgrid
    orig, puz[7][8] = puz[7][8], '1'
    err = puz.Validate()
    if err == nil {
        t.Errorf("no error for duplicate glyph in subgrid 9")
    }
    puz[7][8] = orig
}
