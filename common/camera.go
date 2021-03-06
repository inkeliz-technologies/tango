package common

import (
	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/math32"
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
	// KeyboardZoomerPriority is the priority for the KeyboardZoomerSystem.
	// Priorities determine the order in which the system is updated.
	KeyboardZoomerPriority
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

	tango.Mailbox.ListenMessage(new(CameraMessage), func(msg tango.Message) {
		cammsg, ok := msg.(CameraMessage)
		if !ok {
			return
		}

		if cammsg.Duration > time.Duration(0) {
			currentmsg, ok := cam.longTasks[cammsg.Axis]
			if ok && cammsg.Continue && !currentmsg.Incremental {
				cammsg.Duration = time.Duration((currentmsg.initialValue - cam.getAxis(cammsg.Axis)) / currentmsg.speed)
			}

			cam.longTasks[cammsg.Axis] = &cammsg
			return // because it's handled incrementally
		}

		// Stop with whatever we're doing now
		if _, ok := cam.longTasks[cammsg.Axis]; ok {
			delete(cam.longTasks, cammsg.Axis)
		}

		if cammsg.Incremental {
			cam.moveAxis(cammsg.Axis, cammsg.Value)
		} else {
			cam.moveAxisTo(cammsg.Axis, cammsg.Value)
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
			longTask.initialIncremental = false
			longTask.initialValue = longTask.Value

			longTask.Incremental = true
			longTask.Value -= cam.getAxis(axis)
		}

		// Set speed if needed
		if longTask.speed == 0 {
			longTask.speed = longTask.Value / float32(longTask.Duration.Seconds())
		}

		cam.moveAxis(axis, longTask.speed*dt)

		longTask.Duration -= time.Duration(dt * float32(time.Second))
		if longTask.Duration <= time.Duration(0) {

			// it enforces that the last call will have the exactly position set by `Value` when `incremental = false`.
			if !longTask.initialIncremental {
				cam.moveAxisTo(axis, longTask.initialValue)
			}

			delete(cam.longTasks, axis)
		}
	}

	if cam.tracking.BasicEntity == nil {
		return
	}

	if cam.tracking.SpaceComponent == nil {
		warning("Should be tracking %d but SpaceComponent is nil", cam.tracking.BasicEntity.ID())
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

func (cam *CameraSystem) getAxis(axis CameraAxis) float32 {
	switch axis {
	case XAxis:
		return cam.x
	case YAxis:
		return cam.y
	case ZAxis:
		return cam.z
	case Angle:
		return cam.angle
	default:
		return 0
	}
}

func (cam *CameraSystem) moveAxis(axis CameraAxis, value float32) {
	switch axis {
	case XAxis:
		cam.moveX(value)
	case YAxis:
		cam.moveY(value)
	case ZAxis:
		cam.zoom(value)
	case Angle:
		cam.rotate(value)
	}
}

func (cam *CameraSystem) moveAxisTo(axis CameraAxis, value float32) {
	switch axis {
	case XAxis:
		cam.moveToX(value)
	case YAxis:
		cam.moveToY(value)
	case ZAxis:
		cam.zoomTo(value)
	case Angle:
		cam.rotateTo(value)
	}
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
	Continue    bool

	initialValue, speed float32
	initialIncremental  bool
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

var (
	RotatorSteps90 = []float32{0, 90, 180, 270}
	RotatorSteps45 = []float32{0, 45, 90, 135, 180, 225, 270, 315}
)

// KeyboardScroller is a System that allows for scrolling when certain keys are pressed.
type KeyboardRotator struct {
	RotationSpeed                  float32
	ClockwiseKey, AnticlockwiseKey string
	// DisableHolding prevent holding the key, useful for a stepped rotation like The Sims 1 (90º each press).
	HoldingDisabled bool

	// RotationDuration prevents roughly change of the roration while HoldingDisabled is TRUE
	// It doesn't have effect if HoldingDisabled is FALSE (default)
	RotationDuration time.Duration
	Steps
	camera *CameraSystem
}

// New set default values and camera
func (c *KeyboardRotator) New(w *ecs.World) {
	for _, sys := range w.Systems() {
		switch sys.(type) {
		case *CameraSystem:
			c.camera = sys.(*CameraSystem)
		}
	}

	if c.camera == nil {
		warning("missing camera system in the world")
	}

	c.SetInfinite(true)
}

// Priority implements the ecs.Prioritizer interface.
func (*KeyboardRotator) Priority() int { return KeyboardRotatorPriority }

// Remove does nothing because the KeyboardRotator system has no entities. It implements the
// ecs.System interface.
func (*KeyboardRotator) Remove(ecs.BasicEntity) {}

// Update updates the camera based on keyboard input.
func (c *KeyboardRotator) Update(dt float32) {
	rotation := float32(0)
	if cb := tango.Input.Button(c.ClockwiseKey); cb.JustPressed() || (cb.Down() && !c.HoldingDisabled) {
		rotation += 1
	}

	if ab := tango.Input.Button(c.AnticlockwiseKey); ab.JustPressed() || (ab.Down() && !c.HoldingDisabled) {
		rotation -= 1
	}

	if rotation == 0 {
		return
	}

	rotation *= c.RotationSpeed

	if c.IsStepped() {
		var found bool
		var value float32

		if c.camera.Angle() < 0 || (c.camera.Angle() == 0 && rotation > 0) {
			value, found = c.Steps.Get(math32.Abs(c.camera.Angle()), rotation > 0)
		} else {
			value, found = c.Steps.Get(math32.Abs(c.camera.Angle()), rotation < 0)
		}

		if value == 0 && math32.Abs(c.camera.Angle()) >= 180 {
			value = 360
		}

		if c.camera.Angle() < 0 || (c.camera.Angle() == 0 && rotation > 0) {
			value *= -1
		}

		if !found {
			return
		}

		rotation = value
	}

	tango.Mailbox.Dispatch(CameraMessage{Axis: Angle, Value: rotation, Duration: c.RotationDuration, Incremental: !c.IsStepped()})
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

// KeyboardZoomer is a System that allows for zooming by pressing a keyboard key
type KeyboardZoomer struct {
	ZoomSpeed             float32
	ZoomDuration          time.Duration
	ZoomInKey, ZoomOutKey string
	HoldingDisabled       bool
	Steps

	camera *CameraSystem
}

// New set default values and camera
func (c *KeyboardZoomer) New(w *ecs.World) {
	for _, sys := range w.Systems() {
		switch sys.(type) {
		case *CameraSystem:
			c.camera = sys.(*CameraSystem)
		}
	}

	if c.camera == nil {
		warning("missing camera system in the world")
	}

	if c.ZoomInKey == "" || c.ZoomOutKey == "" {
		c.SetInfinite(true)
	}
}

// Priority implements the ecs.Prioritizer interface.
func (*KeyboardZoomer) Priority() int { return KeyboardZoomerPriority }

// Remove does nothing because KeyboardZoomer has no entities. This implements the
// ecs.System interface.
func (*KeyboardZoomer) Remove(ecs.BasicEntity) {}

// Update zooms the camera in and out based on the movement of the scroll wheel.
func (c *KeyboardZoomer) Update(float32) {
	zoom := float32(0)
	if cb := tango.Input.Button(c.ZoomInKey); cb.JustPressed() || (cb.Down() && !c.HoldingDisabled) {
		zoom += 1
	}

	if ab := tango.Input.Button(c.ZoomOutKey); ab.JustPressed() || (ab.Down() && !c.HoldingDisabled) {
		zoom -= 1
	}

	if zoom == 0 {
		return
	}

	zoom *= c.ZoomSpeed
	if c.IsStepped() {
		value, found := c.Steps.Get(c.camera.Z(), zoom > 0)
		if !found {
			return
		}

		zoom = value
	}

	tango.Mailbox.Dispatch(CameraMessage{Axis: ZAxis, Value: zoom, Duration: c.ZoomDuration, Incremental: !c.IsStepped()})
}

// BindKeyboard sets the vertical and horizontal axes used by the KeyboardScroller.
func (c *KeyboardZoomer) BindKeyboard(zoomInKey, zoomOutKey string) {
	c.ZoomInKey = zoomInKey
	c.ZoomOutKey = zoomOutKey

	if c.ZoomInKey == "" || c.ZoomOutKey == "" {
		c.SetInfinite(true)
	}
}

// KeyboardZoomer creates a new KeyboardZoomer system
func NewKeyboardZoomer(zoomSpeed float32, zoomInKey, zoomOutKey string, disableHolding bool) *KeyboardZoomer {
	return &KeyboardZoomer{
		ZoomSpeed:       zoomSpeed,
		ZoomInKey:       zoomInKey,
		ZoomOutKey:      zoomOutKey,
		HoldingDisabled: disableHolding,
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
	ScrollSpeed              float32
	ScrollLinearAcceleration bool
	// Deprecated use {Top|Right|Bottom|Left}Margin instead
	EdgeMargin                                       float32
	TopMargin, RightMargin, BottomMargin, LeftMargin float32

	currentX, currentY           float32
	maximumX, maximumY           float32
	scrollX, scrollY             float32
	accelerationX, accelerationY float32
	divider                      float32
}

// Priority implements the ecs.Prioritizer interface.
func (*EdgeScroller) Priority() int { return EdgeScrollerPriority }

// Remove does nothing because EdgeScroller has no entities. It implements the ecs.System
// interface.
func (*EdgeScroller) Remove(ecs.BasicEntity) {}

func (c *EdgeScroller) New(_ *ecs.World) {
	if c.TopMargin+c.RightMargin+c.BottomMargin+c.LeftMargin == 0 {
		c.TopMargin, c.RightMargin, c.BottomMargin, c.LeftMargin = c.EdgeMargin, c.EdgeMargin, c.EdgeMargin, c.EdgeMargin
	}
}

// Update moves the camera based on the position of the mouse. If the mouse is on the edge
// of the screen, the camera moves towards that edge.
// TODO: Warning doesn't get the cursor position
func (c *EdgeScroller) Update(dt float32) {
	const cornerSpeed = float32(1.41421356237)

	c.currentX, c.currentY = tango.Input.Mouse.X, tango.Input.Mouse.Y
	c.maximumX, c.maximumY = tango.GameWidth()-c.RightMargin, tango.GameHeight()-c.BottomMargin

	c.scrollX, c.scrollY = float32(0), float32(0)
	c.accelerationX, c.accelerationY = float32(0), float32(0)

	if c.currentY <= c.TopMargin {
		c.scrollY, c.accelerationY = -1, math32.Clamp(c.TopMargin-(c.TopMargin-c.currentY), 1, c.TopMargin)
	}

	if c.currentY >= c.maximumY {
		c.scrollY, c.accelerationY = 1, math32.Clamp(c.BottomMargin-(c.currentY-c.maximumY), 1, c.BottomMargin)
	}

	if c.currentX <= c.LeftMargin {
		c.scrollX, c.accelerationX = -1, math32.Clamp(c.LeftMargin-(c.LeftMargin-c.currentX), 1, c.LeftMargin)
	}

	if c.currentX >= c.maximumX {
		c.scrollX, c.accelerationX = 1, math32.Clamp(c.RightMargin-(c.currentX-c.maximumX), 1, c.RightMargin)
	}

	if c.scrollX == 0 && c.scrollY == 0 {
		return
	}

	if c.scrollX != 0 && c.scrollY != 0 {
		if c.ScrollLinearAcceleration {
			c.divider = math32.Sqrt(2 * ((c.accelerationX + c.accelerationY) / 2))
		} else {
			c.divider = cornerSpeed
		}
	} else {
		c.divider = 1 * (c.accelerationX + c.accelerationY)
	}

	tango.Mailbox.Dispatch(CameraMessage{Axis: XAxis, Value: (c.ScrollSpeed * c.scrollX * dt) / c.divider, Incremental: true})
	tango.Mailbox.Dispatch(CameraMessage{Axis: YAxis, Value: (c.ScrollSpeed * c.scrollY * dt) / c.divider, Incremental: true})
}

func (c *EdgeScroller) SetMargins(top, right, bottom, left float32) {
	c.TopMargin, c.RightMargin, c.BottomMargin, c.LeftMargin = top, right, bottom, left
}

// MouseZoomer is a System that allows for zooming when the scroll wheel is used.
type MouseZoomer struct {
	ZoomSpeed    float32
	ZoomDuration time.Duration
	Steps

	camera *CameraSystem
}

// New set default values and camera
func (c *MouseZoomer) New(w *ecs.World) {
	for _, sys := range w.Systems() {
		switch sys.(type) {
		case *CameraSystem:
			c.camera = sys.(*CameraSystem)
		}
	}

	if c.camera == nil {
		warning("missing camera system in the world")
	}
}

// Priority implements the ecs.Prioritizer interface.
func (*MouseZoomer) Priority() int { return MouseZoomerPriority }

// Remove does nothing because MouseZoomer has no entities. This implements the
// ecs.System interface.
func (*MouseZoomer) Remove(ecs.BasicEntity) {}

// Update zooms the camera in and out based on the movement of the scroll wheel.
func (c *MouseZoomer) Update(float32) {
	scroll := tango.Input.Mouse.ScrollY * c.ZoomSpeed
	if scroll == 0 {
		return
	}

	// if Duration != 0 the c.camera.Z() might be outside of the range of ZoomSteps, so we need to ignore
	if c.IsStepped() {
		value, found := c.Steps.Get(c.camera.Z(), scroll > 0)
		if !found {
			return
		}

		scroll = value
	}

	tango.Mailbox.Dispatch(CameraMessage{Axis: ZAxis, Value: scroll, Duration: c.ZoomDuration, Incremental: !c.IsStepped()})
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

type Steps struct {
	sync.Mutex
	steps    []float32
	infinite bool
}

// SetSteps set the steps slice, it must be ordered
func (s *Steps) SetSteps(steps []float32) {
	// TODO Reorder? e.g if {3, 2, 1} change to {1, 2, 3} ?!
	s.Lock()
	s.steps = steps
	s.Unlock()
}

// SetInifite if true it will return to steps[0] when reach the maximum steps[len(steps)-1]
func (s *Steps) SetInfinite(infinite bool) {
	s.infinite = infinite
}

// IsStepped returns true if there's any steps in use
func (s *Steps) IsStepped() bool {
	return s.steps != nil
}

// Get will returnt the next value (if up is true) or the previous value (if up is false) based on the current.
func (s *Steps) Get(current float32, up bool) (value float32, found bool) {
	s.Lock()

	l := len(s.steps)
	for index, step := range s.steps {
		if step != current {
			continue
		}

		switch {
		case up && l > index+1:
			value, found = s.steps[index+1], true
		case up && s.infinite && l >= index+1:
			value, found = s.steps[0], true
		case !up && index >= 1:
			value, found = s.steps[index-1], true
		case !up && s.infinite && index <= 1:
			value, found = s.steps[l-1], true
		}

		break
	}

	s.Unlock()

	return value, found
}
