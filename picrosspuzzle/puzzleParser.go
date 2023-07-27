package picrosspuzzle

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParsePuzzleFile(filename string) (*Puzzle, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	w_s, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("could not read width line")
	}
	w_s = strings.TrimSpace(w_s)
	width, err := strconv.Atoi(w_s)
	if err != nil {
		return nil, fmt.Errorf("could not parse width %v", w_s)
	}

	h_s, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("could not read height line")
	}
	h_s = strings.TrimSpace(h_s)
	height, err := strconv.Atoi(h_s)
	if err != nil {
		return nil, fmt.Errorf("could not parse height %v", h_s)
	}

	var cols [][]int
	for i := 0; i < width; i++ {
		col_s, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("could not read column %v", i)
		}
		col_s = strings.TrimSpace(col_s)
		var col []int
		for _, c := range strings.Split(col_s, ",") {
			val, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			col = append(col, val)
		}
		cols = append(cols, col)
	}

	var rows [][]int
	for i := 0; i < height; i++ {
		row_s, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("could not read row %v", i)
		}
		row_s = strings.TrimSpace(row_s)
		var row []int
		for _, c := range strings.Split(row_s, ",") {
			val, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			row = append(row, val)
		}
		rows = append(rows, row)
	}

	return NewPuzzle(width, height, rows, cols), nil
}
