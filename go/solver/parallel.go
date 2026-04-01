package solver

func isBoardFull(S *Board) bool {
	for i := range uint8(9) {
		for j := range uint8(9) {
			if S[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// SolveParallel splits the search at the first cell in the MRV order,
// running one goroutine per candidate. The goroutine that finds the
// valid solution wins; others run to completion and are discarded.
func SolveParallel(Si Board) (Board, stats) {
	possibilities := matrixPossibilities(&Si)
	sliceOrder := tableauOrder(&Si)
	if len(sliceOrder) == 0 {
		return Si, stats{}
	}

	first := sliceOrder[0]
	firstPoss := possibilities[first.i][first.j]
	numCandidates := countOnes(firstPoss)

	if numCandidates <= 1 {
		return Solve(Si)
	}

	type result struct {
		board Board
		s     stats
	}
	ch := make(chan result, numCandidates)

	for c := nextCandidate(0, firstPoss); c != 0; c = nextCandidate(c, firstPoss) {
		board := Si
		board[first.i][first.j] = c
		go func(b Board) {
			solved, s := Solve(b)
			ch <- result{solved, s}
		}(board)
	}

	for range numCandidates {
		r := <-ch
		if isBoardFull(&r.board) {
			return r.board, r.s
		}
	}
	return Si, stats{} // impossible
}
