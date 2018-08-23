package ex04_01

import "testing"

type bitDiffTestCase struct {
	A []byte
	B []byte
	E int
}

var bitDiffTests = []bitDiffTestCase{
	{A: []byte{0}, B: []byte{0}, E: 0},
	{A: []byte{0}, B: []byte{1}, E: 1},
	{A: []byte("abc"), B: []byte("bbc"), E: 2},
	{A: []byte{1, 2, 3}, B: []byte{1, 2}, E: -1},
}

func TestBitDiff(t *testing.T) {
	for i := range bitDiffTests {
		a, b := bitDiffTests[i].A, bitDiffTests[i].B
		expected := bitDiffTests[i].E
		got := BitDiff(a, b)
		if got != expected {
			t.Errorf("expected %d, got %d\n", expected, got)
		}
	}
}

var popCountTests = map[byte]uint8{
	byte(0):   uint8(0),
	byte(1):   uint8(1),
	byte(2):   uint8(1),
	byte(3):   uint8(2),
	byte(4):   uint8(1),
	byte(5):   uint8(2),
	byte(6):   uint8(2),
	byte(7):   uint8(3),
	byte(11):  uint8(3),
	byte(15):  uint8(4),
	byte(255): uint8(8),
}

func TestPopCount(t *testing.T) {
	for input, expected := range popCountTests {
		got := PopCount(input)
		if got != expected {
			t.Errorf("expected %d, got %d\n", expected, got)
		}
	}
}
