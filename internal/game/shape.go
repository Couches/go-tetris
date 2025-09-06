package game

import (
	"strings"

	"github.com/couches/go-tetris/internal/mat"
)

type ShapeType byte

const (
	O ShapeType = iota
	I
	T
	J
	L
	S
	Z
)

var shapeMap = map[ShapeType][][]byte{
	O: {
		{1, 1},
		{1, 1},
	},
	I: {
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	T: {
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	},
	J: {
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 0},
	},
	L: {
		{1, 1, 1},
		{1, 0, 0},
		{0, 0, 0},
	},
	S: {
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	},
	Z: {
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	},
}

type Shape struct {
	size int

	grid [][]byte
}

func NewShape(t ShapeType) Shape {
	grid := shapeMap[t]

	return Shape{
		size: len(grid),
		grid: grid,
	}
}

func (s Shape) Rotated(d mat.Direction) [][]byte {
	rotated := make([][]byte, s.size)
	for i := range s.size {
		rotated[i] = make([]byte, s.size)
	}

	switch d {
	case mat.CW:
		for r := range s.size {
			for c := range s.size {
				rotated[c][s.size-1-r] = s.grid[r][c]
			}
		}
	case mat.CCW:
		for r := range s.size {
			for c := range s.size {
				rotated[s.size-1-c][r] = s.grid[r][c]
			}
		}
	}

	return rotated
}

func (s Shape) String() string {
	var sb strings.Builder
	for _, c := range s.grid[0] {
		sb.WriteByte(c + 48)
	}

	for i := range s.size-1 {
		sb.WriteByte(10)
		for _, c := range s.grid[i+1] {
			sb.WriteByte(c + 48)
		}
	}
	return sb.String()
}
