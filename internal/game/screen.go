package game

import (
	"fmt"
	"os"
	"strings"
)

type Screen struct {
	width, height int

	grid [][]string
}

func (s *Screen) Draw() {
	var sb strings.Builder

	if s.height == 0 {
		return
	}

	for _, c := range s.grid[0] {
		sb.WriteString(c)
	}

	if s.height == 1 {
		return
	}

	for i := 1; i < s.height; i++ {
		sb.WriteString("\n")
		for _, c := range s.grid[i] {
			sb.WriteString(c)
		}
	}

	fmt.Print("\033[H")
	os.Stdout.WriteString(sb.String())
}

