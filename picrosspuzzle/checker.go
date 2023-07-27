package picrosspuzzle

import (
	"fmt"
)

func (p *Puzzle) CheckSolution(s *Solution) (valid bool, err error) {
	if p.height != s.height || p.width != s.width {
		return false, fmt.Errorf("width and height do not match; p: %vx%v, s: %vx%v", p.width, p.height, s.width, s.height)
	}

	// Rowwise check
	var runs []int
	var trueRun int
	for row := 0; row < s.height; row++ {
		trueRun = 0
		for col := 0; col < s.width; col++ {
			if s.solution[row][col] {
				trueRun++
				if col == s.width-1 {
					runs = append(runs, trueRun)
				}
			} else if trueRun > 0 {
				runs = append(runs, trueRun)
				trueRun = 0
			}
		}

		if len(runs) != len(p.rows[row]) {
			return false, fmt.Errorf("runs count %v != puzzle clue count %v on row %v", len(runs), len(p.rows[row]), row)
		} else if len(runs) > 0 {
			for i := 0; i < len(runs); i++ {
				if runs[i] != p.rows[row][i] {
					return false, fmt.Errorf("value mismatch run #%v: %v row #%v: %v", i, runs[i], row, p.rows[row][i])
				}
			}
		}
		runs = nil
	}

	// Columnwise check
	runs = nil
	for col := 0; col < s.width; col++ {
		trueRun = 0
		for row := 0; row < s.height; row++ {
			if s.solution[row][col] {
				trueRun++
				if row == s.height-1 {
					runs = append(runs, trueRun)
				}
			} else if trueRun > 0 {
				runs = append(runs, trueRun)
				trueRun = 0
			}
		}

		if len(runs) != len(p.cols[col]) {
			return false, fmt.Errorf("runs count %v != puzzle clue count %v on col %v", len(runs), len(p.cols[col]), col)
		} else if len(runs) > 0 {
			for i := 0; i < len(runs); i++ {
				if runs[i] != p.cols[col][i] {
					return false, fmt.Errorf("value mismatch run #%v: %v col #%v: %v", i, runs[i], col, p.cols[col][i])
				}
			}
		}
		runs = nil
	}

	return true, nil
}
