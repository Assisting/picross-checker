package picrosspuzzle

type Puzzle struct {
	width  int
	height int
	cols   [][]int
	rows   [][]int
}

func NewPuzzle(width int, height int, rows [][]int, cols [][]int) *Puzzle {
	p := Puzzle{
		width:  width,
		height: height,
		rows:   rows,
		cols:   cols,
	}
	return &p
}
