package main

import (
	"log"
	"snakeGame/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := internal.NewGame()

	ebiten.SetWindowSize(internal.ScreenWidth, internal.ScreenHeight)
	ebiten.SetWindowTitle(internal.GameName)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
