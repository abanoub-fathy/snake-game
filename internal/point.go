package internal

import "math/rand"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func NewRandomPoint(xMax, yMax int) *Point {
	return &Point{
		X: rand.Intn(xMax),
		Y: rand.Intn(yMax),
	}
}

func (p *Point) SetX(x int) {
	p.X = x
}

func (p *Point) SetY(y int) {
	p.Y = y
}

func (p *Point) IsLocatedOn(p2 *Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *Point) AddOneStep(stepDirection Direction) {
	switch stepDirection {
	case DirectionRight:
		p.X++
	case DirectionLeft:
		p.X--
	case DirectionUp:
		p.Y--
	case DirectionDown:
		p.Y++
	}
}
