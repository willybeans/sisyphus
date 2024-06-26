package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player *Player
}

func NewGame() *Game {
	g := &Game{}

	g.player = NewPlayer(g)

	return g
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(4, 4)
	// screen.DrawImage(assets.PlayerSprite, op)
	// screen.DrawImage(assets.PlayerSprite, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
