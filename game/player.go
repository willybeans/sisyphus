package game

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"www.github.com/willybeans/sisyphus/assets"
)

const (
	shootCooldown     = time.Millisecond * 500
	rotationPerSecond = math.Pi

	bulletSpawnOffset = 50.0
)

type Player struct {
	game *Game

	position Vector
	// rotation float64
	sprite *ebiten.Image

	// shootCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		game:     game,
		position: pos,
		// rotation: 0,
		sprite: sprite,
		// shootCooldown: NewTimer(shootCooldown),
	}
}

func (p *Player) Update() {
	speed := 5.0
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.position.X += delta.X
	p.position.Y += delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Scale(4, 4)

	// op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
