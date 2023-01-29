package main

import (
	"fmt"
	"time"

	"pkg.amethysts.studio/sudoku-solver-go/solver"
)

func main() {
	sudokus := PlugCLI()

	// multiple sudokus can be given
	for _, sudokuFile := range sudokus {
		PrettyTitle(sudokuFile)

		S := FileHandler(sudokuFile)

		PrettyPrint(S)

		start := time.Now()
		T, stats := solver.Solve(S)
		fmt.Println("Solved in", time.Since(start), "with", stats.Tries, "iterations and", stats.GoingBack, "going back")

		PrettyPrint(T)
	}
}
