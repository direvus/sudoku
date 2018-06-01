package sudoku

import "bytes"

type Mask [Size][Size]bool

func (m *Mask) String() string {
	var buf bytes.Buffer
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if m[i][j] {
				buf.WriteString("✓")
			} else {
				buf.WriteString("✗")
			}
			if j < Size-1 {
				buf.WriteByte(' ')
			} else {
				buf.WriteByte('\n')
			}
		}
	}
	return buf.String()
}

// Equal returns whether two masks contain the same values.
func (a *Mask) Equal(b Mask) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// Fill sets all values of the Mask to 'v'.
func (m *Mask) Fill(v bool) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			m[i][j] = v
		}
	}
}
