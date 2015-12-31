package instruction

type Coord struct {
	X, Y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

type Range struct {
	BL, TR Coord
}

func NewRange(c0, c1 Coord) *Range {
	x0, x1 := c0.X, c1.X
	y0, y1 := c0.Y, c1.Y

	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}

	r := &Range{Coord{x0, y0}, Coord{x1, y1}}
	return r
}

type Instruction struct {
	Cmd   string
	Range r
}

func (r Range) TurnOn() {
	for x := range r {

	}
}

func (r Range) TurnOff() {
}

func (r Range) Toggle() {
}
