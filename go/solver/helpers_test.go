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
