package main

import (
	"encoding"
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
	"strconv"
)

type Point2f struct {
	X, Y float64
}

func Pt2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

func (a Point2f) String() string {
	var (
		sX = strconv.FormatFloat(a.X, 'g', -1, 64)
		sY = strconv.FormatFloat(a.Y, 'g', -1, 64)
	)
	return "(" + sX + "," + sY + ")"
}

var (
	_ encoding.BinaryMarshaler   = Point2f{}
	_ encoding.BinaryUnmarshaler = (*Point2f)(nil)
)

func (a Point2f) MarshalBinary() (data []byte, err error) {
	var (
		uX = math.Float64bits(a.X)
		uY = math.Float64bits(a.Y)
	)
	data = make([]byte, 2*8)
	binary.BigEndian.PutUint64(data[0:], uX)
	binary.BigEndian.PutUint64(data[8:], uY)
	return data, nil
}

func (p *Point2f) UnmarshalBinary(data []byte) error {
	if len(data) < 2*8 {
		return errors.New("insufficient data length")
	}
	var (
		uX = binary.BigEndian.Uint64(data[0:])
		uY = binary.BigEndian.Uint64(data[8:])
	)
	p.X = math.Float64frombits(uX)
	p.Y = math.Float64frombits(uY)
	return nil
}

func (a Point2f) Add(b Point2f) Point2f {
	return Point2f{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point2f) Sub(b Point2f) Point2f {
	return Point2f{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Point2f) MulScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X * scalar,
		Y: a.Y * scalar,
	}
}

func (a Point2f) DivScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X / scalar,
		Y: a.Y / scalar,
	}
}

func randPoint2f(r *rand.Rand) Point2f {
	return Point2f{
		X: r.Float64(),
		Y: r.Float64(),
	}
}

func Distance(a, b Point2f) float64 {
	var (
		dx = a.X - b.X
		dy = a.Y - b.Y
	)
	return math.Sqrt(dx*dx + dy*dy)
}

// A:[a0,a1], B:[b0,b1]
func Intersection(a0, a1, b0, b1 Point2f) (Point2f, bool) {

	var (
		dir0 = a1.Sub(a0)
		dir1 = b1.Sub(b0)
	)

	//считаем уравнения прямых проходящих через отрезки
	var (
		d0Y = -dir0.Y
		d0X = +dir0.X
		d0  = -(d0Y*a0.X + d0X*a0.Y)
	)

	var (
		d1Y = -dir1.Y
		d1X = +dir1.X
		d1  = -(d1Y*b0.X + d1X*b0.Y)
	)

	// подставляем концы отрезков, для выяснения в каких полуплоскотях они
	seg1_line2_start := d1Y*a0.X + d1X*a0.Y + d1
	seg1_line2_end := d1Y*a1.X + d1X*a1.Y + d1

	seg2_line1_start := d0Y*b0.X + d0X*b0.Y + d0
	seg2_line1_end := d0Y*b1.X + d0X*b1.Y + d0

	//если концы одного отрезка имеют один знак, значит он в одной полуплоскости и пересечения нет.
	if (seg1_line2_start*seg1_line2_end >= 0) ||
		(seg2_line1_start*seg2_line1_end >= 0) {
		return Point2f{}, false
	}

	u := seg1_line2_start / (seg1_line2_start - seg1_line2_end)
	p := a0.Add(dir0.MulScalar(u))

	return p, true
}
