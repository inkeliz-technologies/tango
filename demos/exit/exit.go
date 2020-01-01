//+build demo

package main

import (
	"log"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
)

type DefaultScene struct{}

func (*DefaultScene) Preload() {}
func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)
	w.AddSystem(&common.RenderSystem{})
}

func (*DefaultScene) Exit() {
	log.Println("Exit event called; we can do whatever we want now")
	// Here if you want you can prompt the user if they're sure they want to close
	log.Println("Manually closing")
	tango.Exit()
}

func (*DefaultScene) Type() string { return "Game" }

func main() {
	opts := tango.RunOptions{
		Title:               "Exit Demo",
		Width:               1024,
		Height:              640,
		OverrideCloseAction: true,
	}
	tango.Run(opts, &DefaultScene{})
}
