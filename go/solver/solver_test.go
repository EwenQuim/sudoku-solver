package solver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	toSolve := [9][9]uint8{
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{3, 0, 1, 7, 9, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 5, 0, 0, 7, 3, 0, 0},
		{7, 0, 0, 5, 0, 2, 0, 0, 0},
		{0, 0, 8, 0, 1, 0, 2, 0, 0},
		{6, 0, 7, 0, 0, 9, 0, 3, 0},
		{0, 1, 0, 2, 0, 0, 0, 5, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 8},
	}
	result := [9][9]uint8{
		{9, 7, 2, 6, 5, 8, 1, 4, 3},
		{3, 8, 1, 7, 9, 4, 5, 2, 6},
		{5, 4, 6, 3, 2, 1, 9, 8, 7},
		{1, 2, 5, 8, 6, 7, 3, 9, 4},
		{7, 9, 3, 5, 4, 2, 8, 6, 1},
		{4, 6, 8, 9, 1, 3, 2, 7, 5},
		{6, 5, 7, 1, 8, 9, 4, 3, 2},
		{8, 1, 4, 2, 3, 6, 7, 5, 9},
		{2, 3, 9, 4, 7, 5, 6, 1, 8},
	}

	solved, _ := Solve(toSolve)

	if solved != result {
		t.Errorf("Not solved")
	}
}

func BenchmarkSolver(b *testing.B) {
	toSolve := [9][9]uint8{
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{3, 0, 1, 7, 9, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 5, 0, 0, 7, 3, 0, 0},
		{7, 0, 0, 5, 0, 2, 0, 0, 0},
		{0, 0, 8, 0, 1, 0, 2, 0, 0},
		{6, 0, 7, 0, 0, 9, 0, 3, 0},
		{0, 1, 0, 2, 0, 0, 0, 5, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 8},
	}

	for i := 0; i < b.N; i++ {
		_, _ = Solve(toSolve)
	}
}

func BenchmarkSolverHard(b *testing.B) {
	toSolve := [9][9]uint8{
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

	var solved [9][9]uint8
	for i := 0; i < b.N; i++ {
		solved, _ = Solve(toSolve)
	}
	require.NotEqual(b, solved, toSolve)
}

func BenchmarkSolverImpossible(b *testing.B) {
	toSolve := [9][9]uint8{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}

	var solved [9][9]uint8
	for i := 0; i < b.N; i++ {
		solved, _ = Solve(toSolve)
	}
	require.NotEqual(b, solved, toSolve)
}

func begin(Si Board) ([9][9][]uint8, []pos) {
	S := &Si
	// Initialise possibilities, order and digit position
	possibilities := matrixPossibilities(S)
	sliceOrder := tableauOrder(S)
	return possibilities, sliceOrder
}

func BenchmarkBegin(b *testing.B) {
	toSolve := [9][9]uint8{
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

	for i := 0; i < b.N; i++ {
		_, _ = begin(toSolve)
	}
}
