package operations

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radio float64
}

type Triangule struct {
	Base  float64
	Hight float64
}

type Rectangule struct {
	Length float64
	Width  float64
}

func (c *Circle) Area() float64 {
	return c.Radio * 3.1416
}

func (c *Circle) Perimeter() float64 {
	return c.Radio * 2 * 3.1416
}

func (t *Triangule) Area() float64 {

	return (t.Base * t.Hight) / 2
}

func (t *Triangule) Perimeter() float64 {
	return (t.Base) * 3
}

func (r *Rectangule) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangule) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}