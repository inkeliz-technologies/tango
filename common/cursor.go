package common

import (
	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/math32"
)

type CursorSystem struct {
	entities     []mouseEntity
	renderSystem *RenderSystem
	Cursor       *CursorImage
}

type CursorImage struct {
	ecs.BasicEntity
	*RenderComponent
	*SpaceComponent
}

func NewMouseCursor(image Image) *CursorImage {
	return &CursorImage{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: &RenderComponent{
			Hidden:      false,
			Drawable:    NewTextureSingle(image),
			StartZIndex: math32.MaxFloat32,
			StartShader: HUDShader,
			Scale:       tango.Point{X: 0.5, Y: 0.5},
		},
		SpaceComponent: &SpaceComponent{
			Position: tango.CursorPointPos(),
			Width:    float32(image.Width()),
			Height:   float32(image.Height()),
			Rotation: 0,
		},
	}
}

func (m *CursorSystem) New(w *ecs.World) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *RenderSystem:
			m.renderSystem = sys
		}
	}

	m.Cursor = NewMouseCursor(DefaultCursor)
	m.renderSystem.AddByInterface(m.Cursor)
}

func (m *CursorSystem) Remove(basic ecs.BasicEntity) {

}

func (m *CursorSystem) Update(dt float32) {
	if m.Cursor == nil {
		return
	}

	curX, curY := tango.CursorPos()

	m.Cursor.Position.X = math32.Clamp(curX, 1, tango.GameWidth()-1)
	m.Cursor.Position.Y = math32.Clamp(curY, 1, tango.GameHeight()-1)

	tango.SetCursorPosition(m.Cursor.Position.X, m.Cursor.Position.Y)
}

