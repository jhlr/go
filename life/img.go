package life

import (
	"image"
	"image/color"
)

// Pt is a quick way to make a point
func Pt(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

// img implements the image.Image interface
type img struct {
	color  func(bool) color.Color
	model  color.Model
	bounds image.Rectangle
	board  []bool
}

func (i img) ColorModel() color.Model {
	return i.model
}

func (i img) Bounds() image.Rectangle {
	return i.bounds
}

func (i img) At(x, y int) color.Color {
	p := Pt(x, y)
	if !p.In(i.bounds) {
		return i.color(false)
	}
	b := i.Bounds()
	w := b.Dx()
	p = p.Sub(b.Min)
	return i.color(i.board[p.Y*w+p.X])
}

// ColorFunc is a simple callback
// for the SetImage function
func ColorFunc(c color.Color) int {
	r, g, b, a := c.RGBA()
	if a>>15 == 0 {
		return 0
	}
	y := (r + g + b) / 3
	if y >= a/2 {
		return +1
	}
	return -1
}

// GrayFunc is a simple callback
// for the Image function
func GrayFunc(b bool) color.Color {
	if b {
		return color.Gray{255}
	}
	return color.Gray{0}
}

// Image creates a saved state of the full universe
func (u *Universe) Image(cm color.Model, cfoo func(bool) color.Color) image.Image {
	b := u.bounds()
	w := b.Dx()
	h := b.Dy()
	c := u.count % 2
	bools := make([]bool, h*w)
	for j := b.Min.Y; j < b.Max.Y; j++ {
		for i := b.Min.X; i < b.Max.X; i++ {
			p := Pt(i, j)
			cell := u.get(p, c)
			p = p.Sub(b.Min)
			bools[w*p.Y+p.X] = cell
		}
	}
	return img{
		color:  cfoo,
		bounds: b,
		board:  bools,
	}
}

// SetImage will write the contents of the given image
// callback outputs:
// >0 force alive
// =0 keep it how it is
// <0 force dead
func (u *Universe) SetImage(img image.Image, alive func(color.Color) int) {
	b := img.Bounds()
	count := u.count % 2
	for i := b.Min.X; i < b.Max.X; i++ {
		for j := b.Min.Y; j < b.Max.Y; j++ {
			a := alive(img.At(i, j))
			if a != 0 {
				u.set(Pt(i, j), count, a > 0)
			}
		}
	}
	u.Update()
}

// minMax returns the bounds that include
// all the living cells
func (u *Universe) bounds() image.Rectangle {
	u.Update()
	var min, max image.Point
	for p := range u.board {
		min = p
		max = p
		break
	}
	for p := range u.board {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	max.X++
	max.Y++
	return image.Rectangle{
		Min: min,
		Max: max,
	}
}
