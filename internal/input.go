package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) GetDirection() (ebiten.Key, bool) {
	directionKeys := []ebiten.Key{
		ebiten.KeyArrowUp,
		ebiten.KeyArrowDown,
		ebiten.KeyArrowRight,
		ebiten.KeyArrowLeft,
	}

	for _, directionKey := range directionKeys {
		if inpututil.IsKeyJustPressed(directionKey) {
			return directionKey, true
		}
	}

	return 0, false
}
