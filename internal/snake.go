package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	Body      *PointList
	Direction ebiten.Key
	JustAte   bool
}

func NewSnake(body *PointList, direction ebiten.Key) *Snake {
	return &Snake{
		Body:      body,
		Direction: direction,
	}
}

func (s *Snake) GetHead() *Point {
	return s.Body.GetLastPoint()
}

func (s *Snake) GetBodyWithoutHead() *PointList {
	return s.Body.GetPointListWithoutLastPoint()
}

func (s *Snake) CanChangeToNewDirection(newDirection ebiten.Key) bool {
	opposites := map[ebiten.Key]ebiten.Key{
		ebiten.KeyArrowUp:    ebiten.KeyArrowDown,
		ebiten.KeyArrowRight: ebiten.KeyArrowLeft,
		ebiten.KeyArrowDown:  ebiten.KeyArrowUp,
		ebiten.KeyArrowLeft:  ebiten.KeyArrowRight,
	}

	oppositeOfNewDirection, ok := opposites[newDirection]
	if ok && s.Direction != oppositeOfNewDirection {
		return true
	}

	return false
}

func (s *Snake) ChangeDirection(newDirection ebiten.Key) {
	if s.CanChangeToNewDirection(newDirection) {
		s.Direction = newDirection
	}
}

func (s *Snake) IsHeadHitsPoint(point *Point) bool {
	return s.GetHead().IsLocatedOn(point)
}

func (s *Snake) IsHeadHitsBody() bool {
	headPoint := s.GetHead()
	bodyWithoutHead := s.GetBodyWithoutHead()

	return bodyWithoutHead.IncludesPoint(*headPoint)
}

func (s *Snake) Move() {
	headPoint := s.GetHead()
	newHeadPoint := NewPoint(headPoint.X, headPoint.Y)

	switch s.Direction {
	case ebiten.KeyArrowUp:
		newHeadPoint.AddOneStep(DirectionUp)
	case ebiten.KeyArrowDown:
		newHeadPoint.AddOneStep(DirectionDown)
	case ebiten.KeyArrowRight:
		newHeadPoint.AddOneStep(DirectionRight)
	case ebiten.KeyArrowLeft:
		newHeadPoint.AddOneStep(DirectionLeft)
	}

	if s.JustAte {
		s.Body.Push(newHeadPoint)
		s.JustAte = false
	} else {
		s.Body.Shift()
		s.Body.Push(newHeadPoint)
	}
}

func (s *Snake) SetHorizontalCoordinateTo(x int) {
	s.GetHead().SetX(x)
}

func (s *Snake) SetVerticalCoordinateTo(y int) {
	s.GetHead().SetY(y)
}
