package solver

import (
	"sort"
)

// Board represents the sudoku board. uint8 allows a 2x faster processing (and we always have 0 <= 9 <= 255)
type Board = [9][9]uint8

func digitsPossible(S *Board, i uint8, j uint8) uint16 {
	if S[i][j] != 0 { // cell already set
		return 0
	}

	var digits uint16
	for n := uint8(1); n <= 9; n++ {
		if isAvailable(S, i, j, n) {
			digits |= 1 << n
		}
	}

	return digits
}

func matrixPossibilities(S *Board) [9][9]uint16 {
	tab := [9][9]uint16{}

	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			tab[i][j] = digitsPossible(S, i, j)
		}
	}
	return tab
}

type pos struct {
	i uint8
	j uint8
}

// tableauOrder returns the order in which the cells must be processed
func tableauOrder(S *Board) []pos {
	type posWithScores struct {
		pos   pos
		score int // as minimal as possible
	}
	tab := matrixPossibilities(S)

	// compute scores
	listeScores := make([]posWithScores, 0, 81)
	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			if tab[i][j] != 0 {
				listeScores = append(listeScores, posWithScores{
					pos:   pos{i: i, j: j},
					score: 100*countOnes(tab[i][j]) - alignedNeighbors(S, i, j) - squareNeighbors(S, i, j),
				})
			}
		}
	}

	// sort by score
	sort.SliceStable(listeScores, func(a, b int) bool {
		return listeScores[a].score < listeScores[b].score
	})
	liste := make([]pos, 0, len(listeScores))
	for _, p := range listeScores {
		liste = append(liste, p.pos)
	}
	return liste
}

////////////////////////////
//////   Résolution   //////
////////////////////////////

// not really block : only checks the 4 squares that can't be reached with col/row checking
func isAvailableInBloc(S *Board, i uint8, j uint8, n uint8) bool {
	for k := floor3(i); k < floor3(i)+3; k++ {
		if k == i {
			continue
		}
		for l := floor3(j); l < floor3(j)+3; l++ {
			if l != j && S[k][l] == n {
				return false
			}
		}
	}
	return true
}

func isAvailableInLine(S *Board, i uint8, j uint8, n uint8) bool {
	for k := uint8(0); k < 9; k++ {
		if (S[i][k] == n && k != j) || (S[k][j] == n && k != i) {
			return false
		}
	}
	return true
}

func isAvailable(S *Board, i uint8, j uint8, n uint8) bool {
	return isAvailableInLine(S, i, j, n) && isAvailableInBloc(S, i, j, n)
}

func floor3(n uint8) uint8 {
	return 3 * (n / 3)
}

type stats struct {
	Tries     int
	GoingBack int
}

// Solve solves a sudoku and returns the answer
func Solve(Si Board) (Board, stats) {
	S := &Si
	// Initialise possibilities, order and digit position
	possibilities := matrixPossibilities(S)
	sliceOrder := tableauOrder(S)
	maxDigitToFind := uint8(len(sliceOrder))
	var currentCandidate [9][9]uint8
	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			currentCandidate[i][j] = nextCandidate(0, possibilities[i][j])
		}
	}

	// variables
	var rank uint8
	var stats stats

	for rank < maxDigitToFind {
		stats.Tries++

		// Which cell must be processed ?
		n := sliceOrder[rank]
		i := n.i
		j := n.j

		if currentCandidate[i][j] != 0 {
			// There is a digit in the list of possibilities to put here

			if isAvailable(S, i, j, uint8(currentCandidate[i][j])) {
				// digit available -> go forward to next cell
				S[i][j] = uint8(currentCandidate[i][j])
				rank++
			} else {
				// digit already in line, col or 3x3 cell -> try higher digit for same cell
				currentCandidate[i][j] = nextCandidate(currentCandidate[i][j], possibilities[i][j])
			}
		} else {
			// there is no digit possible for this cell

			// first 'reset' the cell state
			S[i][j] = 0
			currentCandidate[i][j] = nextCandidate(0, possibilities[i][j])

			// then go back to previous cell
			rank--
			stats.GoingBack++
			n = sliceOrder[rank]
			i = n.i
			j = n.j
			// previous digit was available but we found it wasn't okay, so increase it
			currentCandidate[i][j] = nextCandidate(currentCandidate[i][j], possibilities[i][j])
		}

	}

	return *S, stats
}
