package common

import (
	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/math32"
	"log"
	"sync"
	"time"
)

const (
	// MouseRotatorPriority is the priority for the MouseRotatorSystem.
	// Priorities determine the order in which the system is updated.
	MouseRotatorPriority = 100 + (iota * 10)
	// MouseZoomerPriority is the priority for he MouseZoomerSystem.
	// Priorities determine the order in which the system is updated.
	MouseZoomerPriority
	// EdgeScrollerPriority is the priority for the EdgeScrollerSystem.
	// Priorities determine the order in which the system is updated.
	EdgeScrollerPriority
	// KeyboardScrollerPriority is the priority for the KeyboardScrollerSystem.
	// Priorities determine the order in which the system is updated.
	KeyboardScrollerPriority
	// KeyboardRotatorPriority is the priority for the KeyboardRotatorSystem.
	// Priorities determine the order in which the system is updated.
	KeyboardRotatorPriority
	// EntityScrollerPriority is the priority for the EntityScrollerSystem.
	// Priorities determine the order in which the system is updated.
	EntityScrollerPriority
)

var (
	// MinZoom is the closest the camera position can be relative to the
	// rendered surface. Smaller numbers of MinZoom allows greater
	// perceived zooming "in".
	MinZoom float32 = 0.25
	// MaxZoom is the farthest the camera position can be relative to the
	// rendered surface. Larger numbers of MaxZoom allows greater
	// perceived zooming "out".
	MaxZoom float32 = 3

	// CameraBounds is the bounding box of the camera
	CameraBounds tango.AABB
)

type cameraEntity struct {
	*ecs.BasicEntity
	*SpaceComponent
}

// CameraSystem is a System that manages the state of the virtual camera. Only
// one CameraSystem can be in a World at a time. If more than one CameraSystem
// is added to the World, it will panic.
type CameraSystem struct {
	x, y, z       float32
	tracking      cameraEntity // The entity that is currently being followed
	trackRotation bool         // Rotate with the entity

	// angle is the angle of the camera, in degrees (not radians!)
	angle float32

	longTasks map[CameraAxis]*CameraMessage
}

// New initializes the CameraSystem.
func (cam *CameraSystem) New(w *ecs.World) {
	num := 0
	for _, sys := range w.Systems() {
		switch sys.(type) {
		case *CameraSystem:
			num++
		}
	}
	if num > 0 { //initalizer is called before added to w.systems
		warning("More than one CameraSystem was added to the World. The RenderSystem adds a CameraSystem if none exist when it's added.")
	}

	if CameraBounds.Max.X == 0 && CameraBounds.Max.Y == 0 {
		CameraBounds.Max = tango.Point{X: tango.GameWidth(), Y: tango.GameHeight()}
	}

	cam.x = CameraBounds.Max.X / 2
	cam.y = CameraBounds.Max.Y / 2
	cam.z = 1

	cam.longTasks = make(map[CameraAxis]*CameraMessage)

	tango.Mailbox.Listen("CameraMessage", func(msg tango.Message) {
		cammsg, ok := msg.(CameraMessage)
		if !ok {
			return
		}

		// Stop with whatever we're doing now
		if _, ok := cam.longTasks[cammsg.Axis]; ok {
			delete(cam.longTasks, cammsg.Axis)
		}

		if cammsg.Duration > time.Duration(0) {
			cam.longTasks[cammsg.Axis] = &cammsg
			return // because it's handled incrementally
		}

		if cammsg.Incremental {
			switch cammsg.Axis {
			case XAxis:
				cam.moveX(cammsg.Value)
			case YAxis:
				cam.moveY(cammsg.Value)
			case ZAxis:
				cam.zoom(cammsg.Value)
			case Angle:
				cam.rotate(cammsg.Value)
			}
		} else {
			switch cammsg.Axis {
			case XAxis:
				cam.moveToX(cammsg.Value)
			case YAxis:
				cam.moveToY(cammsg.Value)
			case ZAxis:
				cam.zoomTo(cammsg.Value)
			case Angle:
				cam.rotateTo(cammsg.Value)
			}
		}
	})

	tango.Mailbox.Dispatch(NewCameraMessage{})
}

// Remove does nothing since the CameraSystem has only one entity, the camera itself.
// This is here to implement the ecs.System interface.
func (cam *CameraSystem) Remove(ecs.BasicEntity) {}

// Update updates the camera. lLong tasks are attempted to update incrementally in batches.
func (cam *CameraSystem) Update(dt float32) {
	for axis, longTask := range cam.longTasks {
		if !longTask.Incremental {
			longTask.Incremental = true

			switch axis {
			case XAxis:
				longTask.Value -= cam.x
			case YAxis:
				longTask.Value -= cam.y
			case ZAxis:
				longTask.Value -= cam.z
			case Angle:
				longTask.Value -= cam.angle
			}
		}

		// Set speed if needed
		if longTask.speed == 0 {
			longTask.speed = longTask.Value / float32(longTask.Duration.Seconds())
		}

		dAxis := longTask.speed * dt
		switch axis {
		case XAxis:
			cam.moveX(dAxis)
		case YAxis:
			cam.moveY(dAxis)
		case ZAxis:
			cam.zoom(dAxis)
		case Angle:
			cam.rotate(dAxis)
		}

		longTask.Duration -= time.Duration(dt)
		if longTask.Duration <= time.Duration(0) {
			delete(cam.longTasks, axis)
		}
	}

	if cam.tracking.BasicEntity == nil {
		return
	}

	if cam.tracking.SpaceComponent == nil {
		log.Println("Should be tracking", cam.tracking.BasicEntity.ID(), "but SpaceComponent is nil")
		cam.tracking.BasicEntity = nil
		return
	}

	cam.centerCam(cam.tracking.SpaceComponent.Position.X+cam.tracking.SpaceComponent.Width/2,
		cam.tracking.SpaceComponent.Position.Y+cam.tracking.SpaceComponent.Height/2,
		cam.z,
	)
	if cam.trackRotation {
		cam.rotateTo(cam.tracking.SpaceComponent.Rotation)
	}
}

// FollowEntity sets the camera to follow the entity with BasicEntity basic
// and SpaceComponent space.
func (cam *CameraSystem) FollowEntity(basic *ecs.BasicEntity, space *SpaceComponent, trackRotation bool) {
	cam.tracking = cameraEntity{basic, space}
	cam.trackRotation = trackRotation
}

// X returns the X-coordinate of the location of the Camera.
func (cam *CameraSystem) X() float32 {
	return cam.x
}

// Y returns the Y-coordinate of the location of the Camera.
func (cam *CameraSystem) Y() float32 {
	return cam.y
}

// Z returns the Z-coordinate of the location of the Camera.
func (cam *CameraSystem) Z() float32 {
	return cam.z
}

// Angle returns the angle (in degrees) at which the Camera is rotated.
func (cam *CameraSystem) Angle() float32 {
	return cam.angle
}

func (cam *CameraSystem) moveX(value float32) {
	if cam.x+(value*tango.GetGlobalScale().X) > CameraBounds.Max.X*tango.GetGlobalScale().X {
		cam.x = CameraBounds.Max.X * tango.GetGlobalScale().X
	} else if cam.x+(value*tango.GetGlobalScale().X) < CameraBounds.Min.X*tango.GetGlobalScale().X {
		cam.x = CameraBounds.Min.X * tango.GetGlobalScale().X
	} else {
		cam.x += value * tango.GetGlobalScale().X
	}
}

func (cam *CameraSystem) moveY(value float32) {
	if cam.y+(value*tango.GetGlobalScale().Y) > CameraBounds.Max.Y*tango.GetGlobalScale().Y {
		cam.y = CameraBounds.Max.Y * tango.GetGlobalScale().Y
	} else if cam.y+(value*tango.GetGlobalScale().Y) < CameraBounds.Min.Y*tango.GetGlobalScale().Y {
		cam.y = CameraBounds.Min.Y * tango.GetGlobalScale().Y
	} else {
		cam.y += value * tango.GetGlobalScale().Y
	}
}

func (cam *CameraSystem) zoom(value float32) {
	cam.zoomTo(cam.z + value)
}

func (cam *CameraSystem) rotate(value float32) {
	cam.rotateTo(cam.angle + value)
}

func (cam *CameraSystem) moveToX(location float32) {
	cam.x = math32.Clamp(location*tango.GetGlobalScale().X, CameraBounds.Min.X*tango.GetGlobalScale().X, CameraBounds.Max.X*tango.GetGlobalScale().X)
}

func (cam *CameraSystem) moveToY(location float32) {
	cam.y = math32.Clamp(location*tango.GetGlobalScale().Y, CameraBounds.Min.Y*tango.GetGlobalScale().Y, CameraBounds.Max.Y*tango.GetGlobalScale().Y)
}

func (cam *CameraSystem) zoomTo(zoomLevel float32) {
	cam.z = math32.Clamp(zoomLevel, MinZoom, MaxZoom)
}

func (cam *CameraSystem) rotateTo(rotation float32) {
	cam.angle = math32.Mod(rotation, 360)
}

func (cam *CameraSystem) centerCam(x, y, z float32) {
	cam.moveToX(x)
	cam.moveToY(y)
	cam.zoomTo(z)
}

// CameraAxis is the axis at which the Camera can/has to move.
type CameraAxis uint8

const (
	// XAxis is the x-axis of the camera
	XAxis CameraAxis = iota
	// YAxis is the y-axis of the camera.
	YAxis
	// ZAxis is the z-axis of the camera.
	ZAxis
	// Angle is the angle the camera is rotated by.
	Angle
)

// CameraMessage is a message that can be sent to the Camera (and other Systemers),
// to indicate movement.
type CameraMessage struct {
	Axis        CameraAxis
	Value       float32
	Incremental bool
	Duration    time.Duration
	speed       float32
}

// Type implements the tango.Message interface.
func (CameraMessage) Type() string {
	return "CameraMessage"
}

// NewCameraMessage is a message that is sent out whenever the camera system changes,
// such as when a new world is created or scenes are switched.
type NewCameraMessage struct{}

// Type implements the tango.Message interface.
func (NewCameraMessage) Type() string {
	return "NewCameraMessage"
}

// KeyboardScroller is a System that allows for scrolling when certain keys are pressed.
type KeyboardScroller struct {
	ScrollSpeed                  float32
	horizontalAxis, verticalAxis string
	keysMu                       sync.RWMutex
}

// Priority implements the ecs.Prioritizer interface.
func (*KeyboardScroller) Priority() int { return KeyboardScrollerPriority }

// Remove does nothing because the KeyboardScroller system has no entities. It implements the
// ecs.System interface.
func (*KeyboardScroller) Remove(ecs.BasicEntity) {}

// Update updates the camera based on keyboard input.
func (c *KeyboardScroller) Update(dt float32) {
	c.keysMu.RLock()
	defer c.keysMu.RUnlock()

	m := tango.Point{
		X: tango.Input.Axis(c.horizontalAxis).Value(),
		Y: tango.Input.Axis(c.verticalAxis).Value(),
	}

	n, _ := m.Normalize()
	if n.X == 0 && n.Y == 0 {
		return
	}

	tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: n.X * c.ScrollSpeed * dt, Incremental: true})
	tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: n.Y * c.ScrollSpeed * dt, Incremental: true})
}

// BindKeyboard sets the vertical and horizontal axes used by the KeyboardScroller.
func (c *KeyboardScroller) BindKeyboard(hori, vert string) {
	c.keysMu.Lock()

	c.verticalAxis = vert
	c.horizontalAxis = hori

	defer c.keysMu.Unlock()
}

// NewKeyboardScroller creates a new KeyboardScroller system using the provided scrollSpeed,
// and horizontal and vertical axes.
func NewKeyboardScroller(scrollSpeed float32, hori, vert string) *KeyboardScroller {
	kbs := &KeyboardScroller{
		ScrollSpeed: scrollSpeed,
	}

	kbs.BindKeyboard(hori, vert)

	return kbs
}

// KeyboardScroller is a System that allows for scrolling when certain keys are pressed.
type KeyboardRotator struct {
	RotationSpeed                  float32
	ClockwiseKey, AnticlockwiseKey string
	// DisableHolding prevent holding the key, useful for a stepped rotation like The Sims 1 (90º each press).
	HoldingDisabled bool

	// RotationAcceleration prevents roughly change of the roration while HoldingDisabled is TRUE
	// It doesn't have effect if HoldingDisabled is FALSE (default)
	// It's not safe to change after defined!
	RotationAcceleration []float32

	remaining float32
	rotation  float32
}

// Priority implements the ecs.Prioritizer interface.
func (*KeyboardRotator) Priority() int { return KeyboardRotatorPriority }

func (c *KeyboardRotator) New(_ *ecs.World) {
	if c.RotationAcceleration == nil {
		c.RotationAcceleration = []float32{c.RotationSpeed}
	}
}

// Remove does nothing because the KeyboardRotator system has no entities. It implements the
// ecs.System interface.
func (*KeyboardRotator) Remove(ecs.BasicEntity) {}

// Update updates the camera based on keyboard input.
func (c *KeyboardRotator) Update(dt float32) {
	if (c.remaining * c.rotation) > 0 {
		l := float32(len(c.RotationAcceleration))

		walk := c.RotationAcceleration[int(math32.Clamp(math32.Ceil(math32.Abs(c.remaining)/(c.RotationSpeed/l))-1, 0, l-1))]
		c.remaining = c.remaining - (walk * c.rotation)

		tango.Mailbox.Dispatch(CameraMessage{Axis: Angle, Value: walk, Incremental: true})
		return
	}

	rotation := float32(0)
	if cb := tango.Input.Button(c.ClockwiseKey); cb.JustPressed() || (cb.Down() && !c.HoldingDisabled) {
		rotation += 1
	}

	if ab := tango.Input.Button(c.AnticlockwiseKey); ab.JustPressed() || (ab.Down() && !c.HoldingDisabled) {
		rotation -= 1
	}

	if rotation != 0 {
		if c.HoldingDisabled {
			c.remaining = rotation * c.RotationSpeed
			c.rotation = rotation
		} else {
			tango.Mailbox.Dispatch(CameraMessage{Axis: Angle, Value: rotation * c.RotationSpeed, Incremental: true})
		}
	}
}

// BindKeyboard sets the vertical and horizontal axes used by the KeyboardScroller.
func (c *KeyboardRotator) BindKeyboard(clockwiseKey, anticlockwiseKey string) {
	c.ClockwiseKey = clockwiseKey
	c.AnticlockwiseKey = anticlockwiseKey
}

// KeyboardRotator creates a new KeyboardRotator system
func NewKeyboardRotator(rotationSpeed float32, clockwiseKey, anticlockwiseKey string, disableHolding bool) *KeyboardRotator {
	return &KeyboardRotator{
		RotationSpeed:    rotationSpeed,
		ClockwiseKey:     clockwiseKey,
		AnticlockwiseKey: anticlockwiseKey,
		HoldingDisabled:  disableHolding,
	}
}

// EntityScroller scrolls the camera to the position of a entity using its space component.
type EntityScroller struct {
	*SpaceComponent
	TrackingBounds tango.AABB
	Rotation       bool
}

// New adjusts CameraBounds to the bounds of EntityScroller.
func (c *EntityScroller) New(*ecs.World) {
	offsetX, offsetY := tango.GameWidth()/2, tango.GameHeight()/2

	CameraBounds.Min.X = c.TrackingBounds.Min.X + (offsetX / tango.GetGlobalScale().X)
	CameraBounds.Min.Y = c.TrackingBounds.Min.Y + (offsetY / tango.GetGlobalScale().Y)

	CameraBounds.Max.X = c.TrackingBounds.Max.X - (offsetX / tango.GetGlobalScale().X)
	CameraBounds.Max.Y = c.TrackingBounds.Max.Y - (offsetY / tango.GetGlobalScale().Y)
}

// Priority implements the ecs.Prioritizer interface.
func (*EntityScroller) Priority() int { return EntityScrollerPriority }

// Remove does nothing because the EntityScroller system has no entities. This implements
// the ecs.System interface.
func (*EntityScroller) Remove(ecs.BasicEntity) {}

// Update moves the camera to the center of the space component.
// Values are automatically clamped to TrackingBounds by the camera.
func (c *EntityScroller) Update(dt float32) {
	if c.SpaceComponent == nil {
		return
	}

	width, height := c.SpaceComponent.Width, c.SpaceComponent.Height

	pos := c.SpaceComponent.Position
	trackToX := pos.X + width/2
	trackToY := pos.Y + height/2

	tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: trackToX, Incremental: false})
	tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: trackToY, Incremental: false})
	if c.Rotation {
		tango.Mailbox.Dispatch(CameraMessage{Axis: Angle, Value: c.SpaceComponent.Rotation, Incremental: false})
	}
}

// EdgeScroller is a System that allows for scrolling when the cursor is near the edges of
// the window.
type EdgeScroller struct {
	ScrollSpeed float32
	EdgeMargin  float32
}

// Priority implements the ecs.Prioritizer interface.
func (*EdgeScroller) Priority() int { return EdgeScrollerPriority }

// Remove does nothing because EdgeScroller has no entities. It implements the ecs.System
// interface.
func (*EdgeScroller) Remove(ecs.BasicEntity) {}

// Update moves the camera based on the position of the mouse. If the mouse is on the edge
// of the screen, the camera moves towards that edge.
// TODO: Warning doesn't get the cursor position
func (c *EdgeScroller) Update(dt float32) {
	curX, curY := tango.CursorPos()
	maxX, maxY := tango.GameWidth(), tango.GameHeight()

	if curX < c.EdgeMargin && curY < c.EdgeMargin {
		s := math32.Sqrt(2)
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: -c.ScrollSpeed * dt / s, Incremental: true})
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: -c.ScrollSpeed * dt / s, Incremental: true})
	} else if curX < c.EdgeMargin && curY > maxY-c.EdgeMargin {
		s := math32.Sqrt(2)
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: -c.ScrollSpeed * dt / s, Incremental: true})
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: c.ScrollSpeed * dt / s, Incremental: true})
	} else if curX > maxX-c.EdgeMargin && curY < c.EdgeMargin {
		s := math32.Sqrt(2)
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: c.ScrollSpeed * dt / s, Incremental: true})
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: -c.ScrollSpeed * dt / s, Incremental: true})
	} else if curX > maxX-c.EdgeMargin && curY > maxY-c.EdgeMargin {
		s := math32.Sqrt(2)
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: c.ScrollSpeed * dt / s, Incremental: true})
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: c.ScrollSpeed * dt / s, Incremental: true})
	} else if curX < c.EdgeMargin {
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: -c.ScrollSpeed * dt, Incremental: true})
	} else if curX > maxX-c.EdgeMargin {
		tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: c.ScrollSpeed * dt, Incremental: true})
	} else if curY < c.EdgeMargin {
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: -c.ScrollSpeed * dt, Incremental: true})
	} else if curY > maxY-c.EdgeMargin {
		tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: c.ScrollSpeed * dt, Incremental: true})
	}
}

// MouseZoomer is a System that allows for zooming when the scroll wheel is used.
type MouseZoomer struct {
	ZoomSpeed float32
}

// Priority implements the ecs.Prioritizer interface.
func (*MouseZoomer) Priority() int { return MouseZoomerPriority }

// Remove does nothing because MouseZoomer has no entities. This implements the
// ecs.System interface.
func (*MouseZoomer) Remove(ecs.BasicEntity) {}

// Update zooms the camera in and out based on the movement of the scroll wheel.
func (c *MouseZoomer) Update(float32) {
	if tango.Input.Mouse.ScrollY != 0 {
		tango.Mailbox.Dispatch(CameraMessage{Axis: ZAxis, Value: tango.Input.Mouse.ScrollY * c.ZoomSpeed, Incremental: true})
	}
}

// MouseRotator is a System that allows for rotating the camera based on pressing
// down the scroll wheel.
type MouseRotator struct {
	// RotationSpeed indicates the speed at which the rotation should happen. This is being used together with the
	// movement by the mouse on the X-axis, to compute the actual rotation.
	RotationSpeed float32

	oldX    float32
	pressed bool
}

// Priority implements the ecs.Prioritizer interface.
func (*MouseRotator) Priority() int { return MouseRotatorPriority }

// Remove does nothing because MouseRotator has no entities. This implements the ecs.System
// interface.
func (*MouseRotator) Remove(ecs.BasicEntity) {}

// Update rotates the camera if the scroll wheel is pressed down.
func (c *MouseRotator) Update(float32) {
	if tango.Input.Mouse.Button == tango.MouseButtonMiddle && tango.Input.Mouse.Action == tango.Press {
		c.pressed = true
	}

	if tango.Input.Mouse.Action == tango.Release {
		c.pressed = false
	}

	if c.pressed {
		tango.Mailbox.Dispatch(CameraMessage{Axis: Angle, Value: (c.oldX - tango.Input.Mouse.X) * -c.RotationSpeed, Incremental: true})
	}

	c.oldX = tango.Input.Mouse.X
}
