package game

import (
	"fmt"

	"github.com/couches/go-tetris/internal/mat"
)

type Piece struct {
	x, y int

	rotation mat.Rotation
	shape    Shape
}

func NewPiece() *Piece {
	return &Piece{}
}

func (p *Piece) SetShape(s Shape) {
	p.shape = s
}

func (p *Piece) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Piece) Overlaps(grid [][]byte) bool {
	gridH := len(grid)
	if gridH == 0 {
		return true
	}

	gridW := len(grid[0])
	if gridW == 0 {
		return true
	}

	shapeEdgeY := p.y+p.shape.size-1
	shapeEdgeX := p.x+p.shape.size-1

	gridEdgeY := gridH-1
	gridEdgeX := gridW-1

	fmt.Println(gridEdgeX)
	fmt.Println(gridEdgeY)

	fmt.Println(shapeEdgeX)
	fmt.Println(shapeEdgeY)

	// entirely off grid?
	if (gridEdgeY < p.y || shapeEdgeY < 0 || gridEdgeX < p.x || shapeEdgeX < 0) {
		return true
	}

	for i, row := range p.shape.grid {
		y := i + p.y

		for j, col := range row {
			if (col == 0) {
				continue
			}

			x := j + p.x

			// out of bounds
			if (y > gridEdgeY || y < 0 || x > gridEdgeX || x < 0) {
				return true
			}

			// overlap
			if (grid[y][x] != 0) {
				return true
			}
		}
	}

	return false
}
