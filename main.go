package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"www.github.com/willybeans/sisyphus/game"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
