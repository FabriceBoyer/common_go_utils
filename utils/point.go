package utils

type Point struct {
	X, Y int
}

func (p *Point) SquaredDistanceFrom(other *Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return float64(dx*dx + dy*dy)
}
