package picrosspuzzle

type Solution struct {
	width    int
	height   int
	solution [][]bool
}

func NewSolution(width int, height int, solution [][]bool) *Solution {
	s := Solution{
		width:    width,
		height:   height,
		solution: solution,
	}
	return &s
}
