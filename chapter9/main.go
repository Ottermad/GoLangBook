package main

import ("fmt"; "math")

// Shape
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// Circle
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r	
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Rectange
type Rectange struct {
	x1, y1, x2, y2 float64
}

func (r *Rectange) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// People
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	//var rx1, ry1 float64 = 0, 0
	//var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5
	r := Rectange{0,0,10,10}
	c := Circle{0, 0, 5}

	fmt.Println(r.area())
	fmt.Println(c.area())
	fmt.Println(totalArea(&c, &r))

	a := new(Android)
	a.Person.Name = "Charlie"
	a.Person.Talk()
	a.Name = "James"
	a.Talk()
}
