//+build demo

package main

import (
	"fmt"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
)

type DefaultScene struct{}

func (*DefaultScene) Preload() {}
func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&InputSystem{})

	tango.Input.RegisterAxis("sideways", tango.AxisKeyPair{tango.KeyA, tango.KeyD})
	tango.Input.RegisterButton("action", tango.KeySpace, tango.KeyEnter)
}

func (*DefaultScene) Type() string { return "Game" }

type inputEntity struct {
	*ecs.BasicEntity
}

type InputSystem struct {
	entities []inputEntity
}

func (c *InputSystem) Add(basic *ecs.BasicEntity) {
	c.entities = append(c.entities, inputEntity{basic})
}

func (c *InputSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *InputSystem) Update(dt float32) {
	if v := tango.Input.Axis("sideways").Value(); v != 0 {
		fmt.Println(v)
	}

	if btn := tango.Input.Button("action"); btn.JustPressed() {
		fmt.Println("DOWN!")
	}
}

func main() {
	opts := tango.RunOptions{
		Title:  "Input Demo",
		Width:  1024,
		Height: 640,
	}

	tango.Run(opts, &DefaultScene{})
}
