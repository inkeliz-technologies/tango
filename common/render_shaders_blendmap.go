package common

import (
	"fmt"

	"github.com/inkeliz-technologies/ecs"
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/gl"
)

type TexturePack struct {
	Fallback *Texture

	RChannel *Texture
	GChannel *Texture
	BChannel *Texture
}

type Blendmap struct {
	*TexturePack

	Map *Texture
}

// Width returns the width of the blendmap.
func (bm Blendmap) Width() float32 {
	return bm.Map.width
}

// Height returns the height of the blendmap.
func (bm Blendmap) Height() float32 {
	return bm.Map.height
}

// Texture returns the OpenGL ID of the blendmap.
func (bm Blendmap) Texture() *gl.Texture {
	return bm.Map.id
}

// View returns the viewport properties of the Texture. The order is Min.X, Min.Y, Max.X, Max.Y.
func (bm Blendmap) View() (float32, float32, float32, float32) {
	return bm.Map.viewport.Min.X, bm.Map.viewport.Min.Y, bm.Map.viewport.Max.X, bm.Map.viewport.Max.Y
}

// Close removes the Texture data from the GPU.
func (bm Blendmap) Close() {
	bm.Map.Close()
}

const (
	blendmapSpriteSize = 20
	// for now we could simply use the default vertex shader.
	blendmapVertexShader   = defaultVertexShader
	blendmapFragmentShader = `
	#ifdef GL_ES
	#define LOWP lowp
	precision mediump float;
	#else
	#define LOWP
	#endif

	varying vec4 var_Color;
	varying vec2 var_TexCoords;

	uniform sampler2D uf_BlendMap;
	uniform sampler2D uf_Fallback;
	uniform sampler2D uf_RChannel;
	uniform sampler2D uf_GChannel;
	uniform sampler2D uf_BChannel;

	uniform vec2 uf_scaleFB;
	uniform vec2 uf_scaleR;
	uniform vec2 uf_scaleG;
	uniform vec2 uf_scaleB;


	vec4 getChan(sampler2D ch, vec2 scale)
	{
		return texture2D(ch, vec2(var_TexCoords.x * scale.x, var_TexCoords.y * scale.y));
	}

	void main(void){
		vec4 mapIdx = texture2D(uf_BlendMap,var_TexCoords);

		vec4 fb = getChan(uf_Fallback, uf_scaleFB) * (1.0 - (mapIdx.r + mapIdx.g + mapIdx.b));
		vec4 r = getChan(uf_RChannel, uf_scaleR) * mapIdx.r;
		vec4 g = getChan(uf_GChannel, uf_scaleG) * mapIdx.g;
		vec4 b = getChan(uf_BChannel, uf_scaleB) * mapIdx.b;

		gl_FragColor=var_Color*(fb+r+g+b);
	}`
)

type blendmapShader struct {
	BatchSize int

	indices     []uint16
	indexBuffer *gl.Buffer
	program     *gl.Program

	vertices                     []float32
	vertexBuffer                 *gl.Buffer
	lastTexture                  *gl.Texture
	lastTexturePack              *TexturePack
	lastRepeating                TextureRepeating
	lastMagFilter, lastMinFilter ZoomFilter

	inPosition  int
	inTexCoords int
	inColor     int

	matrixProjView *gl.UniformLocation
	uf_BlendMap    *gl.UniformLocation
	uf_Fallback    *gl.UniformLocation
	uf_RChannel    *gl.UniformLocation
	uf_GChannel    *gl.UniformLocation
	uf_BChannel    *gl.UniformLocation

	uf_scaleFB *gl.UniformLocation
	uf_scaleR  *gl.UniformLocation
	uf_scaleG  *gl.UniformLocation
	uf_scaleB  *gl.UniformLocation

	projectionMatrix *tango.Matrix
	viewMatrix       *tango.Matrix
	modelMatrix      *tango.Matrix
	cullingMatrix    *tango.Matrix

	camera        *CameraSystem
	cameraEnabled bool

	idx int
}

func (s *blendmapShader) Setup(w *ecs.World) error {
	if s.BatchSize > MaxSprites {
		return fmt.Errorf("%d is greater than the maximum batch size of %d", s.BatchSize, MaxSprites)
	}
	if s.BatchSize <= 0 {
		s.BatchSize = MaxSprites
	}
	// Create the vertex buffer for batching.
	s.vertices = make([]float32, s.BatchSize*blendmapSpriteSize)
	s.vertexBuffer = tango.Gl.CreateBuffer()
	// Create and populate indices buffer. The size of the buffer depends on the batch size.
	// These should never change, so we can just initialize them once here and be done with it.
	numIndicies := s.BatchSize * 6
	s.indices = make([]uint16, numIndicies)
	for i, j := 0, 0; i < numIndicies; i, j = i+6, j+4 {
		s.indices[i+0] = uint16(j + 0)
		s.indices[i+1] = uint16(j + 1)
		s.indices[i+2] = uint16(j + 2)
		s.indices[i+3] = uint16(j + 0)
		s.indices[i+4] = uint16(j + 2)
		s.indices[i+5] = uint16(j + 3)
	}
	var err error
	s.program, err = LoadShader(blendmapVertexShader, blendmapFragmentShader)
	if err != nil {
		return err
	}
	s.indexBuffer = tango.Gl.CreateBuffer()
	tango.Gl.BindBuffer(tango.Gl.ELEMENT_ARRAY_BUFFER, s.indexBuffer)
	tango.Gl.BufferData(tango.Gl.ELEMENT_ARRAY_BUFFER, s.indices, tango.Gl.STATIC_DRAW)

	s.inPosition = tango.Gl.GetAttribLocation(s.program, "in_Position")
	s.inTexCoords = tango.Gl.GetAttribLocation(s.program, "in_TexCoords")
	s.inColor = tango.Gl.GetAttribLocation(s.program, "in_Color")

	s.matrixProjView = tango.Gl.GetUniformLocation(s.program, "matrixProjView")

	s.uf_BlendMap = tango.Gl.GetUniformLocation(s.program, "uf_BlendMap")
	s.uf_Fallback = tango.Gl.GetUniformLocation(s.program, "uf_Fallback")
	s.uf_RChannel = tango.Gl.GetUniformLocation(s.program, "uf_RChannel")
	s.uf_GChannel = tango.Gl.GetUniformLocation(s.program, "uf_GChannel")
	s.uf_BChannel = tango.Gl.GetUniformLocation(s.program, "uf_BChannel")

	s.uf_scaleFB = tango.Gl.GetUniformLocation(s.program, "uf_scaleFB")
	s.uf_scaleR = tango.Gl.GetUniformLocation(s.program, "uf_scaleR")
	s.uf_scaleG = tango.Gl.GetUniformLocation(s.program, "uf_scaleG")
	s.uf_scaleB = tango.Gl.GetUniformLocation(s.program, "uf_scaleB")

	s.projectionMatrix = tango.IdentityMatrix()
	s.viewMatrix = tango.IdentityMatrix()
	s.modelMatrix = tango.IdentityMatrix()
	s.cullingMatrix = tango.IdentityMatrix()

	return nil
}

func (s *blendmapShader) Pre() {
	tango.Gl.Enable(tango.Gl.BLEND)
	tango.Gl.BlendFunc(tango.Gl.SRC_ALPHA, tango.Gl.ONE_MINUS_SRC_ALPHA)
	// Enable shader and buffer, enable attributes in shader
	tango.Gl.UseProgram(s.program)
	tango.Gl.BindBuffer(tango.Gl.ELEMENT_ARRAY_BUFFER, s.indexBuffer)
	tango.Gl.EnableVertexAttribArray(s.inPosition)
	tango.Gl.EnableVertexAttribArray(s.inTexCoords)
	tango.Gl.EnableVertexAttribArray(s.inColor)

	tango.Gl.Uniform1i(s.uf_BlendMap, 0)
	tango.Gl.Uniform1i(s.uf_Fallback, 1)
	tango.Gl.Uniform1i(s.uf_RChannel, 2)
	tango.Gl.Uniform1i(s.uf_GChannel, 3)
	tango.Gl.Uniform1i(s.uf_BChannel, 4)

	// The matrixProjView shader uniform is projection * view.
	// We do the multiplication on the CPU instead of sending each matrix to the shader and letting the GPU do the multiplication,
	// because it's likely faster to do the multiplication client side and send the result over the shader bus than to send two separate
	// buffers over the bus and then do the multiplication on the GPU.
	pv := s.projectionMatrix.Multiply(s.viewMatrix)
	tango.Gl.UniformMatrix3fv(s.matrixProjView, false, pv.Val[:])

	// Since we are batching client side, we only have one VBO, so we can just bind it now and use it for the entire frame.
	tango.Gl.BindBuffer(tango.Gl.ARRAY_BUFFER, s.vertexBuffer)
	tango.Gl.VertexAttribPointer(s.inPosition, 2, tango.Gl.FLOAT, false, 20, 0)
	tango.Gl.VertexAttribPointer(s.inTexCoords, 2, tango.Gl.FLOAT, false, 20, 8)
	tango.Gl.VertexAttribPointer(s.inColor, 4, tango.Gl.UNSIGNED_BYTE, true, 20, 16)
}

func (s *blendmapShader) PrepareCulling() {
	// (Re)initialize the projection matrix.
	s.projectionMatrix.Identity()
	if tango.ScaleOnResize() {
		s.projectionMatrix.Scale(1/(tango.GameWidth()/2), 1/(-tango.GameHeight()/2))
	} else {
		s.projectionMatrix.Scale(1/(tango.CanvasWidth()/(2*tango.CanvasScale())), 1/(-tango.CanvasHeight()/(2*tango.CanvasScale())))
	}
	// (Re)initialize the view matrix
	s.viewMatrix.Identity()
	if s.cameraEnabled {
		s.viewMatrix.Scale(1/s.camera.z, 1/s.camera.z)
		s.viewMatrix.Translate(-s.camera.x, -s.camera.y).Rotate(s.camera.angle)
	} else {
		scaleX, scaleY := s.projectionMatrix.ScaleComponent()
		s.viewMatrix.Translate(-1/scaleX, 1/scaleY)
	}
	s.cullingMatrix.Identity()
	s.cullingMatrix.Multiply(s.projectionMatrix).Multiply(s.viewMatrix)
	s.cullingMatrix.Scale(tango.GetGlobalScale().X, tango.GetGlobalScale().Y)
}

func (s *blendmapShader) ShouldDraw(rc *RenderComponent, sc *SpaceComponent) bool {
	tsc := SpaceComponent{
		Position: sc.Position,
		Width:    rc.Drawable.Width() * rc.Scale.X,
		Height:   rc.Drawable.Height() * rc.Scale.Y,
		Rotation: sc.Rotation,
	}

	c := tsc.Corners()
	c[0].MultiplyMatrixVector(s.cullingMatrix)
	c[1].MultiplyMatrixVector(s.cullingMatrix)
	c[2].MultiplyMatrixVector(s.cullingMatrix)
	c[3].MultiplyMatrixVector(s.cullingMatrix)

	return !((c[0].X < -1 && c[1].X < -1 && c[2].X < -1 && c[3].X < -1) || // All points left of the "viewport"
		(c[0].X > 1 && c[1].X > 1 && c[2].X > 1 && c[3].X > 1) || // All points right of the "viewport"
		(c[0].Y < -1 && c[1].Y < -1 && c[2].Y < -1 && c[3].Y < -1) || // All points above of the "viewport"
		(c[0].Y > 1 && c[1].Y > 1 && c[2].Y > 1 && c[3].Y > 1)) // All points below of the "viewport"
}

func (s *blendmapShader) bindTexturePack(tp *TexturePack) {
	tango.Gl.ActiveTexture(tango.Gl.TEXTURE1)
	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, tp.Fallback.Texture())
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, tango.Gl.REPEAT)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, tango.Gl.REPEAT)

	tango.Gl.ActiveTexture(tango.Gl.TEXTURE2)
	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, tp.RChannel.Texture())
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, tango.Gl.REPEAT)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, tango.Gl.REPEAT)

	tango.Gl.ActiveTexture(tango.Gl.TEXTURE3)
	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, tp.GChannel.Texture())
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, tango.Gl.REPEAT)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, tango.Gl.REPEAT)

	tango.Gl.ActiveTexture(tango.Gl.TEXTURE4)
	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, tp.BChannel.Texture())
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, tango.Gl.REPEAT)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, tango.Gl.REPEAT)

	// always go back to texture 0 since all other shaders might rely on it.
	tango.Gl.ActiveTexture(tango.Gl.TEXTURE0)
}

func (s *blendmapShader) updateScale(bm Blendmap) {
	tango.Gl.Uniform2f(s.uf_scaleFB, bm.Width()/bm.Fallback.Width(), bm.Height()/bm.Fallback.Height())
	tango.Gl.Uniform2f(s.uf_scaleR, bm.Width()/bm.RChannel.Width(), bm.Height()/bm.RChannel.Height())
	tango.Gl.Uniform2f(s.uf_scaleG, bm.Width()/bm.GChannel.Width(), bm.Height()/bm.GChannel.Height())
	tango.Gl.Uniform2f(s.uf_scaleB, bm.Width()/bm.BChannel.Width(), bm.Height()/bm.BChannel.Height())
}

func (s *blendmapShader) Draw(ren *RenderComponent, space *SpaceComponent) {
	bm, ok := ren.Drawable.(Blendmap)
	if !ok {
		panic("only blendmap drawables are supported by blendmap shader.")
	}
	if bm.TexturePack == nil || bm.TexturePack.Fallback == nil {
		panic("No Textures.")
	}

	if s.lastTexturePack != bm.TexturePack {
		s.flush()
		s.bindTexturePack(bm.TexturePack)
		if s.lastTexture == bm.Texture() {
			// if its a different texture we will update the scale with the texture.
			s.updateScale(bm)
		}
		s.lastTexturePack = bm.TexturePack
	}

	if s.lastTexture != ren.Drawable.Texture() {
		s.flush()

		tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, ren.Drawable.Texture())
		s.updateScale(bm)

		s.lastTexture = ren.Drawable.Texture()
	} else if s.idx == len(s.vertices) {
		s.flush()
	}

	if s.lastRepeating != ren.Repeat {
		s.flush()
		var val int
		switch ren.Repeat {
		case NoRepeat:
			val = tango.Gl.CLAMP_TO_EDGE
		case ClampToEdge:
			val = tango.Gl.CLAMP_TO_EDGE
		case ClampToBorder:
			val = tango.Gl.CLAMP_TO_EDGE
		case Repeat:
			val = tango.Gl.REPEAT
		case MirroredRepeat:
			val = tango.Gl.MIRRORED_REPEAT
		}
		tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, val)
		tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, val)
	}

	if s.lastMagFilter != ren.magFilter {
		s.flush()
		var val int
		switch ren.magFilter {
		case FilterNearest:
			val = tango.Gl.NEAREST
		case FilterLinear:
			val = tango.Gl.LINEAR
		}
		tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MAG_FILTER, val)
	}

	if s.lastMinFilter != ren.minFilter {
		s.flush()
		var val int
		switch ren.minFilter {
		case FilterNearest:
			val = tango.Gl.NEAREST
		case FilterLinear:
			val = tango.Gl.LINEAR
		}
		tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MIN_FILTER, val)
	}

	// Update the vertex buffer data.
	s.updateBuffer(ren, space)
	s.idx += 20
}

func (s *blendmapShader) Post() {
	s.flush()
	s.lastTexture = nil
	s.lastTexturePack = nil

	// Cleanup
	tango.Gl.DisableVertexAttribArray(s.inPosition)
	tango.Gl.DisableVertexAttribArray(s.inTexCoords)
	tango.Gl.DisableVertexAttribArray(s.inColor)

	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, nil)
	tango.Gl.BindBuffer(tango.Gl.ARRAY_BUFFER, nil)
	tango.Gl.BindBuffer(tango.Gl.ELEMENT_ARRAY_BUFFER, nil)

	tango.Gl.Disable(tango.Gl.BLEND)
}

func (s *blendmapShader) flush() {
	// If we haven't rendered anything yet, no point in flushing.
	if s.idx == 0 {
		return
	}
	tango.Gl.BufferData(tango.Gl.ARRAY_BUFFER, s.vertices, tango.Gl.STATIC_DRAW)
	// We only want to draw the indicies up to the number of sprites in the current batch.
	count := s.idx / 20 * 6
	tango.Gl.DrawElements(tango.Gl.TRIANGLES, count, tango.Gl.UNSIGNED_SHORT, 0)
	s.idx = 0
	// We need to reset the vertex buffer so that when we start drawing again, we don't accidentally use junk data.
	// The "simpler" way to do this would be to just create a new slice with make(), however that would cause the
	// previous slice to be marked for garbage collection and we'd prefer to keep the GC activity to a minimum.
	for i := range s.vertices {
		s.vertices[i] = 0
	}
}

func (s *blendmapShader) updateBuffer(ren *RenderComponent, space *SpaceComponent) {
	// For backwards compatibility, ren.Buffer is set to the VBO and ren.BufferContent
	// is set to the slice of the vertex buffer for the current sprite. This same slice is
	// populated with vertex data via generateBufferContent.
	ren.Buffer = s.vertexBuffer
	ren.BufferContent = s.vertices[s.idx : s.idx+20]
	s.generateBufferContent(ren, space, ren.BufferContent)
}

func (s *blendmapShader) makeModelMatrix(ren *RenderComponent, space *SpaceComponent) *tango.Matrix {
	// Instead of creating a new model matrix every time, we instead store a global one as a struct member
	// and just reset it for every sprite. This prevents us from allocating a bunch of new Matrix instances in memory
	// ultimately saving on GC activity.
	s.modelMatrix.Identity().Scale(tango.GetGlobalScale().X, tango.GetGlobalScale().Y).Translate(space.Position.X, space.Position.Y)
	if space.Rotation != 0 {
		s.modelMatrix.Rotate(space.Rotation)
	}
	s.modelMatrix.Scale(ren.Scale.X, ren.Scale.Y)
	return s.modelMatrix
}

func (s *blendmapShader) generateBufferContent(ren *RenderComponent, space *SpaceComponent, buffer []float32) bool {
	// We shouldn't use SpaceComponent to get width/height, because this usually already contains the Scale (which
	// is being added elsewhere, so we don't want to over-do it)
	w := ren.Drawable.Width()
	h := ren.Drawable.Height()

	tint := colorToFloat32(ren.Color)

	u, v, u2, v2 := ren.Drawable.View()

	if ren.Repeat != NoRepeat {
		u2 = space.Width / (ren.Drawable.Width() * ren.Scale.X)
		w *= u2
		v2 = space.Width / (ren.Drawable.Height() * ren.Scale.Y)
		h *= v2
	}

	var changed bool

	//setBufferValue(buffer, 0, 0, &changed)
	//setBufferValue(buffer, 1, 0, &changed)
	setBufferValue(buffer, 2, u, &changed)
	setBufferValue(buffer, 3, v, &changed)
	setBufferValue(buffer, 4, tint, &changed)

	setBufferValue(buffer, 5, w, &changed)
	//setBufferValue(buffer, 6, 0, &changed)
	setBufferValue(buffer, 7, u2, &changed)
	setBufferValue(buffer, 8, v, &changed)
	setBufferValue(buffer, 9, tint, &changed)

	setBufferValue(buffer, 10, w, &changed)
	setBufferValue(buffer, 11, h, &changed)
	setBufferValue(buffer, 12, u2, &changed)
	setBufferValue(buffer, 13, v2, &changed)
	setBufferValue(buffer, 14, tint, &changed)

	//setBufferValue(buffer, 15, 0, &changed)
	setBufferValue(buffer, 16, h, &changed)
	setBufferValue(buffer, 17, u, &changed)
	setBufferValue(buffer, 18, v2, &changed)
	setBufferValue(buffer, 19, tint, &changed)

	// Since each sprite in the batch has a different transform, we can't just send the model matrix into
	// the shader and let the GPU take care of it. Instead, we need to multiply the current sprite's model matrix
	// with the position component for each vertex of the current sprite on the CPU, and send the transformed
	// positions to the shader directly.
	modelMatrix := s.makeModelMatrix(ren, space)
	s.multModel(modelMatrix, buffer[:2])
	s.multModel(modelMatrix, buffer[5:7])
	s.multModel(modelMatrix, buffer[10:12])
	s.multModel(modelMatrix, buffer[15:17])
	return changed
}

func (s *blendmapShader) multModel(m *tango.Matrix, v []float32) {
	tmp := tango.MultiplyMatrixVector(m, v)
	v[0] = tmp[0]
	v[1] = tmp[1]
}

func (s *blendmapShader) SetCamera(c *CameraSystem) {
	if s.cameraEnabled {
		s.camera = c
		s.viewMatrix.Identity().Translate(-s.camera.x, -s.camera.y).Rotate(s.camera.angle)
	} else {
		scaleX, scaleY := s.projectionMatrix.ScaleComponent()
		s.viewMatrix.Translate(-1/scaleX, 1/scaleY)
	}
}
