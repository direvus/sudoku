package sudoku

import "bytes"

type Mask [GridSize]bool

func (m *Mask) String() string {
	var buf bytes.Buffer
	for i := 0; i < GridSize; i++ {
		if m[i] {
			buf.WriteString("✓")
		} else {
			buf.WriteString("✗")
		}
		if i % Size == Size-1 {
			buf.WriteByte('\n')
		} else {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

// Equal returns whether two masks contain the same values.
func (a *Mask) Equal(b Mask) bool {
	for i := 0; i < GridSize; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Fill sets all values of the Mask to 'v'.
func (m *Mask) Fill(v bool) {
	for i := 0; i < GridSize; i++ {
		m[i] = v
	}
}

// Count returns the number of cells in the Mask having the given value.
func (m *Mask) Count(v bool) (count int) {
	for i := 0; i < GridSize; i++ {
		if m[i] == v {
			count++
		}
	}
	return
}
