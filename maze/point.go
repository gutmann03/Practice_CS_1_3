package maze

type Point struct {
	X int
	Y int
	p *Point
}

func (p *Point) opposite() *Point {
	if p.X != p.p.X {
		return &Point{X: p.X + (p.X - p.p.X), Y: p.Y, p: p}
	}

	if p.Y != p.p.Y {
		return &Point{X: p.X, Y: p.Y + (p.Y - p.p.Y), p: p}
	}

	return nil
}
