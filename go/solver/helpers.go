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
