package common

import (
	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/gl"
)

type RenderBuffer struct {
	rbo           *gl.RenderBuffer
	width, height int
}

type Framebuffer struct {
	fbo    *gl.FrameBuffer
	oldVP  [4]int32
	isOpen bool
}

type RenderTexture struct {
	tex           *gl.Texture
	width, height float32
	depth         bool
}

func CreateRenderBuffer(width, height int) *RenderBuffer {
	rbuf := &RenderBuffer{
		rbo:    tango.Gl.CreateRenderBuffer(),
		width:  width,
		height: height,
	}
	tango.Gl.BindRenderBuffer(rbuf.rbo)
	tango.Gl.RenderBufferStorage(tango.Gl.RGBA8, width, height)
	tango.Gl.BindRenderBuffer(nil)
	return rbuf
}

func CreateRenderTexture(width, height int, depthBuffer bool) *RenderTexture {
	texBuf := &RenderTexture{
		width:  float32(width),
		height: float32(height),
		tex:    tango.Gl.CreateTexture(),
		depth:  depthBuffer,
	}

	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, texBuf.tex)

	if depthBuffer {
		tango.Gl.TexImage2DEmpty(tango.Gl.TEXTURE_2D, 0, tango.Gl.DEPTH_COMPONENT, width, height, tango.Gl.DEPTH_COMPONENT, tango.Gl.UNSIGNED_BYTE)
	} else {
		tango.Gl.TexImage2DEmpty(tango.Gl.TEXTURE_2D, 0, tango.Gl.RGBA, width, height, tango.Gl.RGBA, tango.Gl.UNSIGNED_BYTE)
	}
	if err := tango.Gl.GetError(); err != 0 {
		panic(err)
	}
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MAG_FILTER, tango.Gl.NEAREST)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MIN_FILTER, tango.Gl.NEAREST)

	return texBuf
}

func CreateFramebuffer() *Framebuffer {
	return &Framebuffer{
		fbo: tango.Gl.CreateFrameBuffer(),
	}
}

func (rb *RenderTexture) Bind() {
	if rb.depth {
		tango.Gl.FrameBufferTexture2D(tango.Gl.FRAMEBUFFER, tango.Gl.DEPTH_ATTACHMENT, tango.Gl.TEXTURE_2D, rb.tex, 0)
	} else {
		tango.Gl.FrameBufferTexture2D(tango.Gl.FRAMEBUFFER, tango.Gl.COLOR_ATTACHMENT0, tango.Gl.TEXTURE_2D, rb.tex, 0)
	}
}

func (t *RenderTexture) Close() {
	tango.Gl.DeleteTexture(t.tex)
}

// Width returns the width of the texture.
func (t *RenderTexture) Width() float32 {
	return t.width
}

// Height returns the height of the texture.
func (t *RenderTexture) Height() float32 {
	return t.height
}

// Texture returns the OpenGL ID of the Texture.
func (t *RenderTexture) Texture() *gl.Texture {
	return t.tex
}

// View returns the viewport properties of the Texture. The order is Min.X, Min.Y, Max.X, Max.Y.
func (t *RenderTexture) View() (float32, float32, float32, float32) {
	return 0, 0, 1, 1
}

func (rb *RenderBuffer) Bind(attachment int) {
	tango.Gl.FrameBufferRenderBuffer(tango.Gl.FRAMEBUFFER, attachment, rb.rbo)
}

func (rb *RenderBuffer) Destroy() {
	tango.Gl.DeleteRenderBuffer(rb.rbo)
}

func (fb *Framebuffer) Open(width, height int) {
	if fb.isOpen {
		return
	}
	tango.Gl.BindFrameBuffer(fb.fbo)
	fb.oldVP = tango.Gl.GetViewport()
	tango.Gl.Viewport(0, 0, width, height)
	fb.isOpen = true
}

func (fb *Framebuffer) Close() {
	if !fb.isOpen {
		return
	}
	tango.Gl.BindFrameBuffer(nil)
	tango.Gl.Viewport(int(fb.oldVP[0]), int(fb.oldVP[1]), int(fb.oldVP[2]), int(fb.oldVP[3]))
	fb.isOpen = false
}

func (fb *Framebuffer) Destroy() {
	tango.Gl.DeleteFrameBuffer(fb.fbo)
}
