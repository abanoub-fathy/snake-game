package internal

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	GameName     = "Snake Game"
	boardRows    = 30
	boardCols    = 30
)

var (
	gameBackgroundColor     = color.RGBA{50, 100, 50, 50}
	gameOverBackgroundColor = color.RGBA{0, 0, 0, 150}
	snakeColor              = color.RGBA{200, 50, 150, 150}
	foodColor               = color.RGBA{200, 200, 50, 150}
	headColor               = color.RGBA{200, 50, 180, 10}
)

type Game struct {
	Input *Input
	Board *Board
}

func NewGame() *Game {
	return &Game{
		Input: NewInput(),
		Board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	// set game background color
	if g.Board.GameOver {
		screen.Fill(gameOverBackgroundColor)
	} else {
		screen.Fill(gameBackgroundColor)
	}

	// draw snake
	width := ScreenHeight / boardRows
	for _, p := range *g.Board.Snake.GetBodyWithoutHead() {
		vector.DrawFilledRect(screen, float32(p.X*width), float32(p.Y*width), float32(width), float32(width), snakeColor, false)
	}

	headPoint := g.Board.Snake.GetHead()
	vector.DrawFilledRect(screen, float32(headPoint.X*width), float32(headPoint.Y*width), float32(width), float32(width), headColor, false)

	if g.Board.GameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Over. Your Score: %d. Press Space to play again...", g.Board.Score))
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Board = NewBoard(boardRows, boardCols)
		}
	} else {
		if g.Board.Food != nil {
			vector.DrawFilledRect(screen, float32(g.Board.Food.Point.X*width), float32(g.Board.Food.Point.Y*width), float32(width), float32(width), foodColor, false)
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.Board.Score))
	}
}

func (g *Game) Update() error {
	return g.Board.Update(g.Input)
}
