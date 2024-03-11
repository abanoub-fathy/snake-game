package internal

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Rows     int
	Cols     int
	Timer    time.Time
	Score    int
	GameOver bool
	Snake    *Snake
	Food     *Food
}

func NewBoard(rows, cols int) *Board {
	board := &Board{
		Rows:     rows,
		Cols:     cols,
		GameOver: false,
		Score:    0,
		Timer:    time.Now(),
	}

	snakeBody := NewPointList(
		NewPoint(0, rows/2),
		NewPoint(1, rows/2),
		NewPoint(2, rows/2),
		NewPoint(3, rows/2),
	)

	board.Snake = NewSnake(snakeBody, ebiten.KeyArrowRight)
	board.PlaceFood()
	return board
}

func (b *Board) BorderLimit(borderDirection Direction) int {
	switch borderDirection {
	case DirectionRight:
		return b.Cols - 1
	case DirectionDown:
		return b.Rows - 1
	default:
		return 0
	}
}

func (b *Board) IsSnakeLeftBorder(borderDirection Direction) bool {
	headPoint := b.Snake.GetHead()

	switch borderDirection {
	case DirectionLeft:
		return headPoint.X < b.BorderLimit(DirectionLeft)
	case DirectionRight:
		return headPoint.X > b.BorderLimit(DirectionRight)
	case DirectionUp:
		return headPoint.Y < b.BorderLimit(DirectionUp)
	default:
		return headPoint.Y > b.BorderLimit(DirectionDown)
	}
}

func (b *Board) PlaceFood() {
	for {
		foodPoint := NewRandomPoint(b.Cols, b.Rows)
		if b.Snake.Body.IncludesPoint(*foodPoint) {
			continue
		}

		b.Food = NewFood(foodPoint)
		break
	}
}

func (b *Board) MoveSnake() error {
	b.Snake.Move()

	if b.Snake.IsHeadHitsBody() {
		b.GameOver = true
		return nil
	}

	if b.IsSnakeLeftBoard() {
		headPoint := b.Snake.GetHead()
		fmt.Println("Head =", *headPoint)

		switch {
		case b.IsSnakeLeftBorder(DirectionLeft):
			b.Snake.SetHorizontalCoordinateTo(b.BorderLimit(DirectionRight))
		case b.IsSnakeLeftBorder(DirectionRight):
			b.Snake.SetHorizontalCoordinateTo(b.BorderLimit(DirectionLeft))
		case b.IsSnakeLeftBorder(DirectionUp):
			b.Snake.SetVerticalCoordinateTo(b.BorderLimit(DirectionDown))
		case b.IsSnakeLeftBorder(DirectionDown):
			b.Snake.SetVerticalCoordinateTo(b.BorderLimit(DirectionUp))
		}
	}

	if b.Snake.IsHeadHitsPoint(b.Food.Point) {
		b.Snake.JustAte = true
		b.Score++

		b.PlaceFood()
	}

	return nil
}

func (b *Board) IsSnakeLeftBoard() bool {
	return b.IsSnakeLeftBorder(DirectionLeft) ||
		b.IsSnakeLeftBorder(DirectionRight) ||
		b.IsSnakeLeftBorder(DirectionUp) ||
		b.IsSnakeLeftBorder(DirectionDown)
}

func (b *Board) Update(input *Input) error {
	if b.GameOver {
		return nil
	}

	interval := time.Millisecond * 100
	if b.Score > 10 {
		interval = time.Millisecond * 80
	} else if b.Score > 20 {
		interval = time.Millisecond * 50
	} else if b.Score > 30 {
		interval = time.Millisecond * 25
	}

	if direction, ok := input.GetDirection(); ok {
		b.Snake.ChangeDirection(direction)
	}

	if time.Since(b.Timer) >= interval {
		if err := b.MoveSnake(); err != nil {
			return err
		}

		b.Timer = time.Now()
	}

	return nil

}
