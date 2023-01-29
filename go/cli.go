package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/profile"
	"pkg.amethysts.studio/sudoku-solver-go/solver"
)

// PlugCLI handles the Command Line Interface
func PlugCLI() []string {

	var profiling = flag.Bool(
		"cpu",
		false,
		"Profiling CPU usage",
	)
	flag.Usage = func() {
		color.Set(color.Bold)
		fmt.Fprintf(os.Stderr, "Usage of: ")
		color.Blue(os.Args[0])
		color.Unset()

		flag.PrintDefaults()

		fmt.Println("  args\n\tPath to a sudoku file. Can be a list of files or something like `data/*`")

	}

	flag.Parse()

	if *profiling {
		defer profile.Start().Stop()
	}

	return flag.Args()
}

// PrettyPrint prints matrix
func PrettyPrint(S solver.Board) {
	for i := 0; i < 9; i++ {
		fmt.Println(S[i])
	}
	fmt.Println()
}

// PrettyTitle prints sudoku title
func PrettyTitle(sudokuFile string) {
	color.Set(color.Bold)
	fmt.Print("File ")
	color.Blue(sudokuFile)
	color.Unset()
}
