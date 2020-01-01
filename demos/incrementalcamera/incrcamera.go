//+build demo

package main

import (
	"image/color"
	"time"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
	"github.com/inkeliz-technologies/tango/demos/demoutils"
)

type DefaultScene struct{}

var (
	worldWidth  int = 800
	worldHeight int = 800
)

func (*DefaultScene) Preload() {}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)
	w.AddSystem(&common.RenderSystem{})

	demoutils.NewBackground(w, worldWidth, worldHeight, color.RGBA{102, 153, 0, 255}, color.RGBA{102, 173, 0, 255})

	// Center camera if GlobalScale is Setup
	tango.Mailbox.Dispatch(common.CameraMessage{Axis: common.XAxis,
		Value:       float32(worldWidth) / 2,
		Incremental: false})
	tango.Mailbox.Dispatch(common.CameraMessage{Axis: common.YAxis,
		Value:       float32(worldHeight) / 2,
		Incremental: false})

	// We issue one camera zoom command at the start, but it takes a while to process because we set a duration
	tango.Mailbox.Dispatch(common.CameraMessage{
		Axis:        common.ZAxis,
		Value:       3, // so zooming out a lot
		Incremental: true,
		Duration:    time.Second * 5,
	})
}

func (*DefaultScene) Type() string { return "Game" }

func main() {
	opts := tango.RunOptions{
		Title:  "IncrementalCamera Demo",
		Width:  worldWidth,
		Height: worldHeight,
	}
	tango.Run(opts, &DefaultScene{})
}
