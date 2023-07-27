package picrosspuzzle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseSolutionFile(filename string) (*Solution, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	width := 0
	var solution [][]bool
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.New("could not read  line")
		}
		cells := strings.Split(strings.TrimSpace(line), "")
		var row []bool
		for i := 0; i < len(cells); i++ {
			if cells[i] == "_" {
				row = append(row, false)
			} else if cells[i] == "x" {
				row = append(row, true)
			} else {
				return nil, fmt.Errorf("found non-solution chracter %v", cells[i])
			}
		}

		if width == 0 {
			width = len(row)
		} else if width != len(row) {
			return nil, fmt.Errorf("column count changed on row %v", len(solution))
		}
		solution = append(solution, row)
	}

	s := NewSolution(width, len(solution), solution)
	return s, nil
}
