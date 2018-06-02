package sudoku

import "testing"

func TestMaskString(t *testing.T) {
	m := Mask{
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false}}
	expect := (
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n")
	result := m.String()
	if expect != result {
		t.Errorf("invalid string output, expected\n%v\n\ngot\n%v", expect, result)
	}
	m = Mask{
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true}}
	expect = (
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n" +
		"✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓\n")
	result = m.String()
	if expect != result {
		t.Errorf("invalid string output, expected\n%v\n\ngot\n%v", expect, result)
	}
	m = Mask{
		{false, false, false, false, false, false, false, false, false},
		{false, false,  true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false,  true},
		{false, false, false, false, false, false, false, false, false}}
	expect = (
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✓ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✓\n" +
		"✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗ ✗\n")
	result = m.String()
	if expect != result {
		t.Errorf("invalid string output, expected\n%v\n\ngot\n%v", expect, result)
	}
}

func TestMaskEqual(t *testing.T) {
	a := Mask{
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true}}
	b := Mask{
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, true, true, true}}
	if a.Equal(b) {
		t.Errorf("incorrect positive return from Equal")
	}
	if b.Equal(a) {
		t.Errorf("incorrect positive return from Equal")
	}

	b = Mask{
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true}}
	if !a.Equal(b) {
		t.Errorf("incorrect negative return from Equal")
	}
	if !b.Equal(a) {
		t.Errorf("incorrect negative return from Equal")
	}
}

func TestMaskFill(t *testing.T) {
	var m Mask
	m.Fill(true)
	expect := Mask{
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true}}
	if !m.Equal(expect) {
		t.Errorf("incorrect result from Fill: expected all true values, got\n%v", m.String())
	}
	m.Fill(false)
	expect = Mask{
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false}}
	if !m.Equal(expect) {
		t.Errorf("incorrect result from Fill: expected all false values, got\n%v", m.String())
	}
}

func TestMaskCount(t *testing.T) {
	var m Mask
	count := m.Count(false)
	expect := Size*Size
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
	count = m.Count(true)
	expect = 0
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
	m.Fill(true)
	count = m.Count(true)
	expect = Size*Size
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
	count = m.Count(false)
	expect = 0
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
	m = Mask{
		{ true, false,  true, false, false,  true, false,  true, false},
		{false,  true, false, false,  true, false,  true,  true, false},
		{ true, false,  true,  true, false,  true, false,  true,  true},
		{false,  true, false, false,  true,  true, false,  true, false},
		{ true, false,  true,  true, false, false, false,  true, false},
		{ true, false,  true, false,  true, false,  true, false,  true},
		{false,  true, false, false,  true,  true, false,  true, false},
		{ true, false,  true, false, false,  true, false,  true, false},
		{false,  true,  true, false,  true,  true, false,  true,  true}}
	count = m.Count(true)
	expect = 41
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
	count = m.Count(false)
	expect = 40
	if count != expect {
		t.Errorf("incorrect result from Count: expected %v, got %v", expect, count)
	}
}
