package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Player struct {
	*components.Camera
	*components.Position
}

func (p *Player) Update(w engine.World) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += 5
	}
}

var _ engine.SystemUpdater = (*Player)(nil)
