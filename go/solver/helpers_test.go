package solver

import "testing"

func TestNeighbors(t *testing.T) {
	toSolve := &[9][9]uint8{
		{1, 0, 0, 0, 0, 7, 0, 9, 0},
		{0, 3, 0, 0, 2, 0, 0, 0, 8},
		{0, 0, 9, 6, 0, 0, 5, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 0},
		{0, 1, 0, 0, 8, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 4, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 4, 1, 0, 0, 0, 0, 0, 7},
		{0, 0, 7, 0, 0, 0, 3, 0, 0},
	}

	neighbors := alignedNeighbors(toSolve, 8, 8)
	if neighbors != 5 {
		t.Errorf("Wrong number of neighbors")
	}
}

func TestSquareNeighbors(t *testing.T) {
	toSolve := &[9][9]uint8{
		{1, 0, 0, 0, 0, 7, 0, 9, 0},
		{0, 3, 0, 0, 2, 0, 0, 0, 8},
		{0, 0, 9, 6, 0, 0, 5, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 0},
		{0, 1, 0, 0, 8, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 4, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 4, 1, 0, 0, 0, 0, 0, 7},
		{0, 0, 7, 0, 0, 0, 3, 0, 0},
	}

	neighbors := squareNeighbors(toSolve, 8, 8)
	if neighbors != 3 {
		t.Errorf("Wrong number of neighbors")
	}

	neighbors = squareNeighbors(toSolve, 7, 7)
	if neighbors != 3 {
		t.Errorf("Wrong number of neighbors")
	}
}

func TestCountOnes(t *testing.T) {
	tests := []struct {
		input  uint16
		output int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{63, 6}, // 2^6 - 1 = 111_111
		{64, 1}, // 2^6 = 1_000_000
	}

	for _, test := range tests {
		if countOnes(test.input) != test.output {
			t.Errorf("countOnes(%d) = %d, want %d", test.input, countOnes(test.input), test.output)
		}
	}
}

func TestNextCandidate(t *testing.T) {
	tests := []struct {
		input      uint8
		candidates uint16
		output     uint8
	}{
		{0, 0b010110, 1},
		{1, 0b010110, 2},
		{2, 0b010110, 4},
		{3, 0b010110, 4},
		{4, 0b010110, 0},
		{5, 0b010110, 0},

		{0, 0b1100100100, 2},
	}

	for _, test := range tests {
		if nextCandidate(test.input, test.candidates) != test.output {
			t.Errorf("nextCandidate(%d) = %d, want %d", test.input, nextCandidate(test.input, test.candidates), test.output)
		}
	}
}

func BenchmarkNexCandidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextCandidate(0, 0b010110)
	}
}
