package main

import (
	"fmt"
	"strings"

	"github.com/couches/go-tetris/internal/game"
)

func main() {
	/*
	g := game.NewGame()
	gl := gameLoop.New(5, g.Update)

	gl.Start()

	for {}
	*/

	b := [][]byte {
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
	}

	s := game.NewShape(game.S)
	p := game.NewPiece()
	p.SetShape(s)
	p.SetPosition(1, 3)

	fmt.Println(GridString(b))
	fmt.Println()
	fmt.Println(s.String())
	fmt.Println(p.Overlaps(b))
}

func GridString(grid [][]byte) string {
	var sb strings.Builder
	for _, c := range grid[0] {
		sb.WriteByte(c + 48)
	}

	for i := range len(grid)-1 {
		sb.WriteByte(10)
		for _, c := range grid[i+1] {
			if c == 0 {
				sb.WriteRune('.')
				continue
			}
			sb.WriteByte(c + 48)
		}
	}
	return sb.String()
}
