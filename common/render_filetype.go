package common

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	// imported to decode jpegs and upload them to the GPU.
	_ "image/jpeg"
	// imported to decode .pngs and upload them to the GPU.
	_ "image/png"
	// imported to decode .gifs and uppload them to the GPU.
	_ "image/gif"
	"io"

	// these are for svg support

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	"github.com/inkeliz-technologies/tango"
	"github.com/inkeliz-technologies/tango/gl"
)

// TextureResource is the resource used by the RenderSystem. It uses .jpg, .gif, and .png images
type TextureResource struct {
	Texture *gl.Texture
	Width   float32
	Height  float32
	url     string
}

// URL is the file path of the TextureResource
func (t TextureResource) URL() string {
	return t.url
}

type imageLoader struct {
	images map[string]TextureResource
}

func (i *imageLoader) Load(url string, data io.Reader) error {
	if getExt(url) == ".svg" {
		icon, err := oksvg.ReadIconStream(data, oksvg.WarnErrorMode)
		if err != nil {
			return err
		}
		w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
		img := image.NewRGBA(image.Rect(0, 0, w, h))
		gv := rasterx.NewScannerGV(w, h, img, img.Bounds())
		r := rasterx.NewDasher(w, h, gv)
		icon.Draw(r, 1.0)
		b := img.Bounds()
		newm := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(newm, newm.Bounds(), img, b.Min, draw.Src)
		i.images[url] = NewTextureResource(&ImageObject{newm})
	} else {
		img, _, err := image.Decode(data)
		if err != nil {
			return err
		}
		b := img.Bounds()
		newm := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(newm, newm.Bounds(), img, b.Min, draw.Src)
		i.images[url] = NewTextureResource(&ImageObject{newm})
	}

	return nil
}

func (i *imageLoader) Unload(url string) error {
	delete(i.images, url)
	return nil
}

func (i *imageLoader) Resource(url string) (tango.Resource, error) {
	texture, ok := i.images[url]
	if !ok {
		return nil, fmt.Errorf("resource not loaded by `FileLoader`: %q", url)
	}

	return texture, nil
}

// Image holds data and properties of an .jpg, .gif, or .png file
type Image interface {
	image.Image
	Data() interface{}
	Width() int
	Height() int
}

// UploadTexture sends the image to the GPU, to be kept in GPU RAM
func UploadTexture(img Image) (id *gl.Texture) {
	if tango.Headless() {
		return id
	}

	if img.Data() == nil {
		panic("Texture image data is nil.")
	}

	id = tango.Gl.CreateTexture()
	tango.Gl.BindTexture(tango.Gl.TEXTURE_2D, id)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_S, tango.Gl.CLAMP_TO_EDGE)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_WRAP_T, tango.Gl.CLAMP_TO_EDGE)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MIN_FILTER, tango.Gl.LINEAR)
	tango.Gl.TexParameteri(tango.Gl.TEXTURE_2D, tango.Gl.TEXTURE_MAG_FILTER, tango.Gl.NEAREST)
	tango.Gl.TexImage2D(tango.Gl.TEXTURE_2D, 0, tango.Gl.RGBA, tango.Gl.RGBA, tango.Gl.UNSIGNED_BYTE, img.Data())

	return id
}

// NewTextureResource sends any image.Image to the GPU and returns a `TextureResource` for easy access
func NewTextureResource(img image.Image) TextureResource {
	obj, ok := img.(Image)
	if !ok {
		obj = NewImageObject(img)
	}

	return TextureResource{Texture: UploadTexture(obj), Width: float32(obj.Width()), Height: float32(obj.Height())}
}

// NewTextureSingle sends any image.Image to the GPU and returns a `Texture` with a viewport for single-sprite images
func NewTextureSingle(img image.Image) Texture {
	obj, ok := img.(Image)
	if !ok {
		obj = NewImageObject(img)
	}

	return Texture{id: UploadTexture(obj), width: float32(obj.Width()), height: float32(obj.Height()), viewport: tango.AABB{Max: tango.Point{X: 1.0, Y: 1.0}}}
}

// ImageToNRGBA takes a given `image.Image` and converts it into an `image.NRGBA`. Especially useful when transforming
// image.Uniform to something usable by `tango`.
func ImageToNRGBA(img image.Image, width, height int) (nrgba *image.NRGBA) {
	nrgba = image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(nrgba, nrgba.Rect, img, image.Point{X: 0, Y: 0}, draw.Src)

	return nrgba
}

// ImageObject is a pure Go implementation of a `Drawable`
type ImageObject struct {
	data *image.NRGBA
}

// NewImageObject creates a new ImageObject given any image.Image reference
func NewImageObject(img image.Image) *ImageObject {
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		size := img.Bounds().Size()
		nrgba = ImageToNRGBA(img, size.X, size.Y)
	}

	return NewImageObjectNRGBA(nrgba)
}

// NewImageObjectNRGBA creates a new ImageObject given image.NRGBA reference
func NewImageObjectNRGBA(nrgba *image.NRGBA) *ImageObject {
	return &ImageObject{data: nrgba}
}

// ColorModel implements image.Image
func (i *ImageObject) ColorModel() color.Model {
	return i.data.ColorModel()
}

// Bounds implements image.Image
func (i *ImageObject) Bounds() image.Rectangle {
	return i.data.Bounds()
}

// At implements image.Image
func (i *ImageObject) At(x, y int) color.Color {
	return i.data.At(x, y)
}

// Data returns the entire image.NRGBA object
func (i *ImageObject) Data() interface{} {
	return i.data
}

// Width returns the maximum X coordinate of the image
func (i *ImageObject) Width() int {
	return i.data.Rect.Max.X
}

// Height returns the maximum Y coordinate of the image
func (i *ImageObject) Height() int {
	return i.data.Rect.Max.Y
}

// LoadedSprite loads the texture-reference from `tango.Files`, and wraps it in a `*Texture`.
// This method is intended for image-files which represent entire sprites.
func LoadedSprite(url string) (*Texture, error) {
	res, err := tango.Files.Resource(url)
	if err != nil {
		return nil, err
	}

	img, ok := res.(TextureResource)
	if !ok {
		return nil, fmt.Errorf("resource not of type `TextureResource`: %s", url)
	}

	return &Texture{img.Texture, img.Width, img.Height, tango.AABB{Max: tango.Point{X: 1.0, Y: 1.0}}}, nil
}

// Texture represents a texture loaded in the GPU RAM (by using OpenGL), which defined dimensions and viewport
type Texture struct {
	id       *gl.Texture
	width    float32
	height   float32
	viewport tango.AABB
}

// Width returns the width of the texture.
func (t Texture) Width() float32 {
	return t.width
}

// Height returns the height of the texture.
func (t Texture) Height() float32 {
	return t.height
}

// Texture returns the OpenGL ID of the Texture.
func (t Texture) Texture() *gl.Texture {
	return t.id
}

// View returns the viewport properties of the Texture. The order is Min.X, Min.Y, Max.X, Max.Y.
func (t Texture) View() (float32, float32, float32, float32) {
	return t.viewport.Min.X, t.viewport.Min.Y, t.viewport.Max.X, t.viewport.Max.Y
}

// Close removes the Texture data from the GPU.
func (t Texture) Close() {
	if tango.Headless() {
		return
	}

	tango.Gl.DeleteTexture(t.id)
}

func init() {
	tango.Files.Register(".jpg", &imageLoader{images: make(map[string]TextureResource)})
	tango.Files.Register(".png", &imageLoader{images: make(map[string]TextureResource)})
	tango.Files.Register(".gif", &imageLoader{images: make(map[string]TextureResource)})
	tango.Files.Register(".svg", &imageLoader{images: make(map[string]TextureResource)})
}
