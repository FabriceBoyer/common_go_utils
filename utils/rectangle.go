package utils

type Rect struct {
	X, Y, W, H int
}

func (rect *Rect) IsValid() bool {
	return rect.X != 0 && rect.Y != 0 && rect.W != 0 && rect.H != 0
}

func (rect *Rect) Top() int {
	if rect.H > 0 {
		return rect.Y
	}
	return rect.Y + rect.H
}

func (rect *Rect) Bottom() int {
	if rect.H > 0 {
		return rect.Y + rect.H
	}
	return rect.Y
}

func (rect *Rect) Left() int {
	if rect.W > 0 {
		return rect.X
	}
	return rect.X + rect.W
}

func (rect *Rect) Right() int {
	if rect.W > 0 {
		return rect.X + rect.W
	}
	return rect.X
}

func (rect *Rect) Reset() {
	rect.X = 0
	rect.Y = 0
	rect.W = 0
	rect.H = 0
}

func (rect *Rect) Intersects(other *Rect) bool {
	return other.Left() < rect.Right() && other.Right() > rect.Left() &&
		other.Top() < rect.Bottom() && other.Bottom() > rect.Top()
}

func (rect *Rect) BBCenter() *Point {
	return &Point{rect.X + rect.W/2, rect.Y + rect.H/2}
}
