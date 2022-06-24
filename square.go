package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{
		x: s.a,
		y: s.a,
	}
}

func (s Square) Area() uint {
	return uint(s.a) * uint(s.a)
}

func (s Square) Perimeter() uint {
	return uint(s.a) + uint(s.a)
}
