package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Debug struct {
	*component.Debug
}

func (d *Debug) Draw(_ engine.World, screen *ebiten.Image) {
	if !d.Debug.Enabled {
		return
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %.2f FPS: %.2f",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(),
	))
}