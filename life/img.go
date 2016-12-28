package life

import (
	"image"
	"image/color"
)

type Position image.Point

func P(x, y int) Position {
	return Position{X: x, Y: y}
}

type img struct {
	min, max image.Point
	board    [][]bool
}

func (i img) ColorModel() color.Model {
	return color.GrayModel
}

func (i img) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: i.min,
		Max: i.max,
	}
}

func (i img) At(x, y int) color.Color {
	vx := x >= i.min.X && x < i.max.X
	vy := y >= i.min.Y && y < i.max.Y
	x = x - i.min.X
	y = y - i.min.Y
	if vx && vy && i.board[y][x] {
		return color.Gray{255}
	}
	return color.Gray{0}
}

func (u *Universe) Image() image.Image {
	min, max := u.minMax()
	max.X++
	max.Y++
	return img{
		min:   image.Point(min),
		max:   image.Point(max),
		board: u.Bools(),
	}
}

func (u *Universe) SetImage(img image.Image) {
	b := img.Bounds()
	count := u.count % 2
	for i := b.Min.X; i < b.Max.X; i++ {
		for j := b.Min.Y; j < b.Max.Y; j++ {
			c := img.At(i, j)
			c = color.GrayModel.Convert(c)
			ui8 := c.(color.Gray).Y
			u.set(P(i, j), count, ui8 >= 128)
		}
	}
	u.trim()
}

func (u *Universe) minMax() (Position, Position) {
	u.trim()
	var min, max Position
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
	return min, max
}

func (u *Universe) Bools() [][]bool {
	min, max := u.minMax()
	res := make([][]bool, max.Y-min.Y+1)
	c := u.count % 2
	for j := min.Y; j <= max.Y; j++ {
		res[j-min.Y] = make([]bool, max.X-min.X+1)
		for i := min.X; i <= max.X; i++ {
			p := P(i, j)
			res[j-min.Y][i-min.X] = u.get(p, c)
		}
	}
	return res
}
