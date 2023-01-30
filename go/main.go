package main

import (
	"fmt"
	"time"

	"pkg.amethysts.studio/sudoku-solver-go/solver"
)

func main() {
	sudokus, silent := PlugCLI()

	if len(sudokus) == 0 {
		fmt.Println("No sudoku file given")
		return
	}

	if !silent {
		fmt.Println("Solving", len(sudokus), "sudokus...")
	}

	// multiple sudokus can be given
	for _, sudokuFile := range sudokus {
		S := FileHandler(sudokuFile)

		if !silent {
			PrettyTitle(sudokuFile)
			PrettyPrint(S)
		}

		start := time.Now()
		T, stats := solver.Solve(S)

		if !silent {
			fmt.Println("Solved in", time.Since(start), "with", stats.Tries, "iterations and", stats.GoingBack, "going back")
			PrettyPrint(T)
		}
	}
}
