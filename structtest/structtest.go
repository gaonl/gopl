package main

import "fmt"

type People struct {
	Name string
	Age int
}

type Animal struct {
	Name string
	Kind string
}

type World struct {
	People
	Animal
}

type Point struct {
	X, Y int
}
func (p *Point) Add(point *Point) {
	p.X += point.X
	p.Y += point.Y
}

type ColorPoint struct {
	Point
	Color int
}
func (p *ColorPoint) Add(point *ColorPoint) {
	p.X += point.X
	p.Y += point.Y
}

func main(){
	/*w := World{}
	w.People.Name = "abc"
	w.Age = 100
	w.Kind = "dog"
	w.Animal.Name = "hash chee"

	fmt.Println(w)

	template.ParseFiles().Par()*/
	pointB := &Point{1,2}
	point := &Point{5,6}
	(point).Add(pointB)
	fmt.Println(point)

	ca := &ColorPoint{
		Point: Point{1, 2},
		Color: 1,
	}

	cb := &ColorPoint{
		Point: Point{1, 2},
		Color: 1,
	}

	ca.Add(cb)
	fmt.Println(ca)
}
