package initialser

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"strings"
	"image"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/fixed"
	"image/jpeg"
	"bytes"
	"image/png"
	"errors"
	"github.com/golang/freetype"
	"io"
)

type Drawer struct {
	face   font.Face
	font   *truetype.Font
	dpi    float64
	avatar *Avatar
	scale  fixed.Int26_6
}
//NewDrawer create a Drawer
func NewDrawer(avatar *Avatar) (*Drawer, error) {
	trueFont, err := _font_cache.get(avatar.Font)
	if err != nil {
		return nil, err
	}
	d := &Drawer{
		font: trueFont,
		dpi:  72.0,
		avatar:avatar,
		scale:fixed.Int26_6(avatar.FontSize << 6),
	}
	d.face = truetype.NewFace(d.font, &truetype.Options{
		Size:    float64(avatar.FontSize),
		DPI:     d.dpi,
		Hinting: font.HintingFull,
	})
	return d, nil
}

//Draw draw
func (d *Drawer) Draw() (image.Image, error) {
	bg, err := hexToRGBA(d.avatar.Background);
	if err != nil {
		return nil, err
	}
	fc, err := hexToRGBA(d.avatar.Color);
	if err != nil {
		return nil, err
	}
	dst := image.NewRGBA(image.Rect(0, 0, d.avatar.Size, d.avatar.Size))
	draw.Draw(dst, dst.Bounds(), bg, image.ZP, draw.Src)
	fd := &font.Drawer{
		Dst: dst,
		Src: fc,
		Face: truetype.NewFace(d.font, &truetype.Options{
			Size:    float64(d.avatar.FontSize),
			DPI:     d.dpi,
			Hinting: font.HintingNone,
		}),
	}
	dot, err := d.center()
	if err != nil {
		return nil, err
	}
	fd.Dot = dot
	fd.DrawString(string(d.avatar.initial))
	return dst, nil
}

//center cal draw text center
func (d *Drawer)center() (fixed.Point26_6, error) {
	var gb truetype.GlyphBuf
	err := gb.Load(d.font, d.scale, d.font.Index(d.avatar.initial), font.HintingFull)
	if err != nil {
		return fixed.Point26_6{}, err
	}
	bounds := gb.Bounds
	dp := freetype.Pt(d.avatar.Size, d.avatar.Size).Sub(bounds.Max.Sub(bounds.Min));
	y := dp.Y / 2 + bounds.Max.Y
	x := dp.X / 2 - bounds.Min.X
	return fixed.Point26_6{X:x, Y:y}, nil
}
//DrawToBytes draw image data to []byte
func (d *Drawer) DrawToBytes(encoding ...string) ([]byte, error) {
	var buf bytes.Buffer
	err := d.DrawToWriter(&buf, encoding...);
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
//DrawToWriter draw image data to writer
func (d *Drawer) DrawToWriter(w io.Writer, encoding ...string) error {
	encode := "png"
	if len(encoding) > 0 {
		encode = encoding[0];
	}
	encode = strings.ToLower(encode)
	m, err := d.Draw()
	if err != nil {
		return err
	}
	switch encode {
	case "jpeg", "jpg":
		err = jpeg.Encode(w, m, nil)
	case "png":
		err = png.Encode(w, m)
	default:
		err = errors.New("not support ext " + encode)
	}
	return err
}
