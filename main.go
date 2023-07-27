package main

import (
	"fmt"
	"os"
	picross "picross/picrosspuzzle"
	"strings"
)

func main() {
	puzzleFilename := "example-puzzle.txt"
	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {
		puzzleFilename = strings.TrimSpace(os.Args[1])
	}

	solutionFilename := "example-solution.txt"
	if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" {
		solutionFilename = strings.TrimSpace(os.Args[2])
	}

	puzzle, err := picross.ParsePuzzleFile(puzzleFilename)
	if err != nil {
		panic(err)
	}

	solution, err := picross.ParseSolutionFile(solutionFilename)
	if err != nil {
		panic(err)
	}

	result, err := puzzle.CheckSolution(solution)
	if result {
		fmt.Println("Valid solution.")
	} else {
		fmt.Println("INVALID solution!")
		fmt.Println(err)
	}
}
