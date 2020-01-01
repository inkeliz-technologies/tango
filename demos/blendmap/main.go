package main

import (
	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
)

type DefaultScene struct{}

func (*DefaultScene) Preload() {
	err := tango.Files.Load("blendmap.png", "grass.png", "mud.png", "path.png", "flowers.png")
	if err != nil {
		panic(err)
	}
}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)

	rs := new(common.RenderSystem)
	w.AddSystem(rs)

	pack := new(common.TexturePack)
	pack.RChannel, _ = common.LoadedSprite("flowers.png")
	pack.GChannel, _ = common.LoadedSprite("mud.png")
	pack.BChannel, _ = common.LoadedSprite("path.png")
	pack.Fallback, _ = common.LoadedSprite("grass.png")

	blendMap, _ := common.LoadedSprite("blendmap.png")

	ent := &sampleEntity{
		ecs.NewBasic(),
		common.SpaceComponent{},
		common.RenderComponent{
			Scale:    tango.Point{0.75, 0.75},
			Drawable: common.Blendmap{pack, blendMap},
		},
	}

	rs.Add(&ent.BasicEntity, &ent.RenderComponent, &ent.SpaceComponent)
}

func (*DefaultScene) Type() string { return "Game" }

type sampleEntity struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func main() {
	opts := tango.RunOptions{
		Title:  "Blendmap Demo",
		Width:  800,
		Height: 800,
	}

	tango.Run(opts, &DefaultScene{})
}
