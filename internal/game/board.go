package game

type Board struct {
	width, height byte

	grid [][]byte
}

func NewBoard(width, height byte) *Board {
	grid := make([][]byte, height)
	for i := range height {
		grid[i] = make([]byte, width)
	}

	return &Board{
		width: width,
		height: height,

		grid: grid,
	}
}
