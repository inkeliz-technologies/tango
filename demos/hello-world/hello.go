//+build demo

package main

import (
	"image/color"
	"log"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
)

type DefaultScene struct{}

type Guy struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
}

func (*DefaultScene) Preload() {
	tango.Files.Load("icon.png")
}

func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)

	w.AddSystem(&common.RenderSystem{})

	// Retrieve a texture
	texture, err := common.LoadedSprite("icon.png")
	if err != nil {
		log.Println(err)
	}

	// Create an entity
	guy := Guy{BasicEntity: ecs.NewBasic()}

	// Initialize the components, set scale to 8x
	guy.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    tango.Point{8, 8},
	}
	guy.SpaceComponent = common.SpaceComponent{
		Position: tango.Point{0, 0},
		Width:    texture.Width() * guy.RenderComponent.Scale.X,
		Height:   texture.Height() * guy.RenderComponent.Scale.Y,
	}

	// Add it to appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&guy.BasicEntity, &guy.RenderComponent, &guy.SpaceComponent)
		}
	}
}

func (*DefaultScene) Type() string { return "GameWorld" }

func main() {
	opts := tango.RunOptions{
		Title:  "Hello World Demo",
		Width:  1024,
		Height: 640,
	}
	tango.Run(opts, &DefaultScene{})
}
