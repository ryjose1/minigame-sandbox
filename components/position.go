package components

type Position struct {
	x      int
	y      int
	width  int
	height int
}

func NewPosition(x int, y int, width int, height int) *Position {
	return &Position{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (p *Position) X() int {
	return p.x
}

func (p *Position) SetX(x int) {
	p.x = x
}

func (p *Position) Y() int {
	return p.y
}

func (p *Position) SetY(y int) {
	p.y = y
}

func (p *Position) Width() int {
	return p.width
}

func (p *Position) Height() int {
	return p.height
}
