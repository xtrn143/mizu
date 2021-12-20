package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Spin struct {
	*component.Root
	*component.Angle
	*component.Spin
}

func (s *Spin) Update(_ engine.World) {
	if s.Root.Root {
		return
	}

	s.Angle.Deg += s.Spin.Speed
}