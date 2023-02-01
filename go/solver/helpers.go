package solver

func alignedNeighbors(S *Board, i, j uint8) int {
	count := 0
	for k := uint8(0); k < 9; k++ {
		if j != k && S[i][k] != 0 {
			count++
		}
		if i != k && S[k][j] != 0 {
			count++
		}
	}

	return count
}

func squareNeighbors(S *Board, i, j uint8) int {
	count := 0
	for k := floor3(uint8(i)); k < floor3(uint8(i))+3; k++ {
		for l := floor3(uint8(j)); l < floor3(uint8(j))+3; l++ {
			if S[k][l] != 0 && (k != i || l != j) {
				count++
			}
		}
	}

	return count
}

func countOnes(n uint16) int {
	var count int
	for n > 0 {
		count++
		n &= n - 1
	}
	return count
}

func nextCandidate(from uint8 /* 0 to 9 */, possibilities uint16) uint8 /* 1 to 9 */ {
	for i := from + 1; i <= 9; i++ {
		if possibilities&(1<<i) != 0 {
			return i
		}
	}
	return 0
}
