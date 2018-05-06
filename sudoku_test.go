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
        t.Errorf("Read failed to notice malformed input.")
    }

    // Invalid puzzle input (too few lines)
    err = puz.Read([]byte(
        "1 _ 3 _ _ 6 _ 8 _\n" +
        "_ 5 _ _ 8 _ 1 2 _\n"))
    if err == nil {
        t.Errorf("Read failed to notice insufficient lines.")
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
        t.Errorf("Read failed to notice junk characters around input.")
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
        t.Errorf("Read failed to notice invalid glyph in row 6.")
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
        t.Errorf("Read failed to notice insufficient glyphs in row 6.")
    }
}
