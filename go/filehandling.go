package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"pkg.amethysts.studio/sudoku-solver-go/solver"
)

// FileHandler opens a file by its name and returns string
func FileHandler(fileName string) solver.Board {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	S := solver.Board{}
	for i, eachline := range fileTextLines {
		for j, character := range eachline {
			c := string(character)
			if c == "." {
				S[i][j] = 0
			} else {
				convert, _ := strconv.Atoi(c)
				S[i][j] = uint8(convert)
			}
		}

	}
	return S
}
