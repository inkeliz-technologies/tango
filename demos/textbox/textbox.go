//+build demo

package main

import (
	"image/color"
	"sync"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/common"
)

type DefaultScene struct{}

type MyLabel struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*DefaultScene) Preload() {
	err := tango.Files.Load("Roboto-Regular.ttf")
	if err != nil {
		panic(err)
	}
	tango.Input.RegisterButton("backspace", tango.KeyBackspace)
	tango.Input.RegisterButton("enter", tango.KeyEnter)
}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(u tango.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&TypingSystem{})

	fnt := &common.Font{
		URL:  "Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	label1 := MyLabel{BasicEntity: ecs.NewBasic()}
	label1.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Start Typing to add text!",
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&label1.BasicEntity, &label1.RenderComponent, &label1.SpaceComponent)
		}
	}
}

func (*DefaultScene) Type() string { return "Game" }

type TypingSystem struct {
	label             MyLabel
	runeLock          sync.Mutex
	runesToAdd        []rune
	timeSinceDeletion float32
}

func (t *TypingSystem) New(w *ecs.World) {
	fnt := &common.Font{
		URL:  "Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	t.label = MyLabel{BasicEntity: ecs.NewBasic()}
	t.label.SpaceComponent.Position.Set(0, 75)
	t.label.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "",
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.label.BasicEntity, &t.label.RenderComponent, &t.label.SpaceComponent)
		}
	}

	tango.Mailbox.Listen("TextMessage", func(msg tango.Message) {
		m, ok := msg.(tango.TextMessage)
		if !ok {
			return
		}
		t.runeLock.Lock()
		t.runesToAdd = append(t.runesToAdd, m.Char)
		t.runeLock.Unlock()
	})
}

func (*TypingSystem) Remove(ecs.BasicEntity) {}

func (t *TypingSystem) Update(dt float32) {
	t.timeSinceDeletion += dt
	str := ""
	txt := t.label.Drawable.(common.Text)
	if tango.Input.Button("backspace").Down() && len(txt.Text) > 0 {
		if t.timeSinceDeletion > 0.2 {
			t.timeSinceDeletion = 0
			txt.Text = txt.Text[:len(txt.Text)-1]
			t.label.Drawable = txt
		}
		return
	}
	t.runeLock.Lock()
	if len(t.runesToAdd) != 0 {
		str = string(t.runesToAdd)
	}
	t.runesToAdd = make([]rune, 0)
	t.runeLock.Unlock()
	txt.Text += str
	if tango.Input.Button("enter").JustPressed() {
		txt.Text += "\n"
	}
	t.label.Drawable = txt
}

func main() {
	opts := tango.RunOptions{
		Title:  "Typing Demo",
		Width:  800,
		Height: 800,
	}
	tango.Run(opts, &DefaultScene{})
}
