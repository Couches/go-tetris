package mat

type Rotation byte

const (
	N Rotation = iota
	E
	S
	W
)

type Direction byte

const (
	CW Direction = iota
	CCW
)

func (r Rotation) Rotated(d Direction) Rotation {
	switch d {
	case CW:
		return (r + 1) % 4
	case CCW:
		return (r + 3) % 4
	default:
		return r
	}
}
