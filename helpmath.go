package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"

	"math/rand/v2"
)

// POINTS
func POINTONCIRC(radius, angl float64, cntr Vector2) Vector2 {
	angleInRadians := angl * math.Pi / 180.0
	x := radius*math.Cos(angleInRadians) + cntr.X
	y := radius*math.Sin(angleInRadians) + cntr.Y
	return Vector2{x, y}
}
func POINTINRECXY(p sdl.Point, x, y, width, height int32) bool {
	return p.X >= x && p.X <= x+width && p.Y >= y && p.Y <= y+height
}
func POINTINREC(p sdl.Point, r sdl.Rect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}
func POINTFINREC(p sdl.FPoint, r sdl.Rect) bool {
	return p.X >= float32(r.X) && p.X <= float32(r.X+r.W) && p.Y >= float32(r.Y) && p.Y <= float32(r.Y+r.H)
}
func ANGL2POINTS(v1, v2 Vector2) float64 {
	deltaX := v2.X - v1.X
	deltaY := v2.Y - v1.Y
	radians := math.Atan2(deltaY, deltaX)
	degrees := (radians * 180) / math.Pi
	for degrees >= 360 {
		degrees -= 360
	}
	for degrees < 0 {
		degrees += 360
	}
	return degrees
}
func REMǁPOINT(slice []sdl.Point, s int) []sdl.Point {
	return append(slice[:s], slice[s+1:]...)
}

// RECS
func RECCENTER(r sdl.Rect) sdl.Point {
	return sdl.Point{r.X + r.W/2, r.Y + r.H/2}
}
func RECCENTERV2(r sdl.Rect) Vector2 {
	return Vector2{float64(r.X), float64(r.Y)}
}
func RECSINTERSECT(r1, r2 sdl.Rect) bool {
	return r1.X < r2.X+r2.W &&
		r1.X+r1.W > r2.X &&
		r1.Y < r2.Y+r2.H &&
		r1.Y+r1.H > r2.Y
}
func ROCNTR(r sdl.Rect) *sdl.Point {
	return &sdl.Point{r.W / 2, r.H / 2}
}

/*
	func imfromtile(t TILESHEET, r sdl.Rect) IM {
		i := IM{}
		i.tex = t.tex
		i.r = r
		return i
	}

	func remANIM(slice []ANIM, s int) []ANIM {
		return append(slice[:s], slice[s+1:]...)
	}

	func objSORT(o []OBJ) []OBJ {
		//sort.Slice(o, func(i, j int) bool { return o[i].zi > o[j].zi })
		return o
	}
*/

func orbit(cx, cy, radius, angle float64) (float64, float64) {
	x := cx + radius*math.Cos(angle)
	y := cy + radius*math.Sin(angle)
	return x, y
}
func rotate(cnt Vector2, rad, angl float64) sdl.FPoint { //BETTER VERSION
	radians := angl * math.Pi / 180.0
	newX := cnt.X + rad*math.Cos(float64(radians))
	newY := cnt.Y + rad*math.Sin(float64(radians))
	return sdl.FPoint{float32(newX), float32(newY)}
}
func velXYnewpoint(originalPoint sdl.FPoint, xVelocity, yVelocity, maxSpeed, angle float64) (float32, float32) {

	angleRad := angle * math.Pi / 180.0

	newXVelocity := xVelocity*math.Cos(angleRad) - yVelocity*math.Sin(angleRad)
	newYVelocity := xVelocity*math.Sin(angleRad) + yVelocity*math.Cos(angleRad)

	newSpeed := math.Sqrt(newXVelocity*newXVelocity + newYVelocity*newYVelocity)
	if newSpeed > maxSpeed {
		scaleFactor := maxSpeed / newSpeed
		newXVelocity *= scaleFactor
		newYVelocity *= scaleFactor
	}

	return float32(newXVelocity), float32(newYVelocity)
}

func resizerec(r sdl.Rect, scale int32) sdl.Rect {
	return sdl.Rect{RECCENTER(r).X - ((r.W * scale) / 2), RECCENTER(r).Y - ((r.H * scale) / 2), r.W * scale, r.H * scale}
}

type Vector2 struct {
	X, Y float64
}

func sqcntrF(cnt sdl.FPoint, w float32) sdl.FRect {
	return sdl.FRect{cnt.X - w/2, cnt.Y - w/2, w, w}
}

func vel2pointsVECFPOINT(p1 Vector2, p2 sdl.FPoint, maxspeed float32) (xvel float32, yvel float32) {
	xvel = p2.X - float32(p1.X)
	yvel = p2.Y - float32(p1.Y)

	distance := float32(math.Sqrt(float64(xvel*xvel + yvel*yvel)))

	if distance > maxspeed {
		scale := maxspeed / distance
		xvel *= scale
		yvel *= scale
	}

	return xvel, yvel
}

func v2tofpoint(v Vector2) *sdl.FPoint {
	return &sdl.FPoint{float32(v.X), float32(v.Y)}
}
func v2tofpoint2(v Vector2) sdl.FPoint {
	return sdl.FPoint{float32(v.X), float32(v.Y)}
}
func fpoint2v2(p sdl.FPoint) Vector2 {
	return Vector2{float64(p.X), float64(p.Y)}
}
func point2v2(p sdl.Point) Vector2 {
	return Vector2{float64(p.X), float64(p.Y)}
}
func rec2frec(r sdl.Rect) sdl.FRect {
	return sdl.FRect{float32(r.X), float32(r.Y), float32(r.W), float32(r.H)}
}
func FREC2REC(r sdl.FRect) sdl.Rect {
	return sdl.Rect{int32(r.X), int32(r.Y), int32(r.W), int32(r.H)}
}
func FREC2RECǁPOINTER(r sdl.FRect) *sdl.Rect {
	return &sdl.Rect{int32(r.X), int32(r.Y), int32(r.W), int32(r.H)}
}
func normalizeSpeed(x, y, maxspd float64) (float64, float64) {
	v := Vector2{x, y}
	v2 := Normalize(v)
	newx := v2.X * maxspd
	newy := v2.Y * maxspd
	return newx, newy
}

func Normalize(v Vector2) Vector2 {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if magnitude == 0 {
		return v // Avoid division by zero
	}
	return Vector2{v.X / magnitude, v.Y / magnitude}
}

func recResizeHeight(w, h, newW int32) int32 {
	return (newW * h) / w
}
func recResizeWidth(w, h, newH int32) int32 {
	return (newH * w) / h
}
func origin(r sdl.Rect) sdl.Point {
	return sdl.Point{r.W / 2, r.H / 2}
}

func originF(r sdl.FRect) *sdl.FPoint {
	return &sdl.FPoint{r.W / 2, r.H / 2}
}
func origin2int32toF(r sdl.Rect) *sdl.FPoint {
	return &sdl.FPoint{float32(r.W / 2), float32(r.H / 2)}
}
func checkpointtri(p sdl.Point, tri []sdl.Point) bool {
	as_x := p.X - tri[0].X
	as_y := p.Y - tri[0].Y
	s_ab := (tri[1].X-tri[0].X)*as_y-(tri[1].Y-tri[0].Y)*as_x > 0
	if (tri[2].X-tri[0].X)*as_y-(tri[2].Y-tri[0].Y)*as_x > 0 == s_ab {
		return false
	}
	if (tri[2].X-tri[1].X)*(p.Y-tri[1].Y)-(tri[2].Y-tri[1].Y)*(p.X-tri[1].X) > 0 != s_ab {
		return false
	}
	return true
}

func ChangeAlpha(col []uint8, alpha uint8) []uint8 {
	col[3] = alpha
	return col
}

func RecsFIntersect(r1, r2 sdl.FRect) bool {
	return r1.X < r2.X+r2.W &&
		r1.X+r1.W > r2.X &&
		r1.Y < r2.Y+r2.H &&
		r1.Y+r1.H > r2.Y
}

func Point2FPoint(p sdl.Point) sdl.FPoint {
	return sdl.FPoint{float32(p.X), float32(p.Y)}
}
func ColorSliceToSDLColor(color []uint8) sdl.Color {
	return sdl.Color{color[0], color[1], color[2], color[3]}
}
func ColorSDLtoSlice(c sdl.Color) []uint8 {
	return []uint8{c.R, c.G, c.B, c.A}
}
func Abs(n float32) float32 {
	return float32(math.Abs(float64(n)))
}

func AbsDiff(n, n2 float32) float32 {
	return float32(math.Abs(float64(n) - float64(n2)))
}

func FLIPCOIN() bool {
	onoff := false
	choose := RINT(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}
func Roll6() int {
	return RINT(1, 7)
}
func Roll12() int {
	return RINT(1, 13)
}
func Roll18() int {
	return RINT(1, 19)
}
func Roll24() int {
	return RINT(1, 25)
}
func Roll30() int {
	return RINT(1, 31)
}
func Roll36() int {
	return RINT(1, 37)
}
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RUINT8(min2, max2 uint8) uint8 {
	min := int(min2)
	max := int(max2)
	return uint8(min + rand.IntN(max-min))
}
func RI32(min, max int) int32 {
	return int32(min + rand.IntN(max-min))
}
func RF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
func RF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func StF32(t float32) string {
	return fmt.Sprint(t)
}
func StF32NODEC(t float32) string {
	return fmt.Sprintf("%.0f", t)
}
func StINT(t int) string {
	return fmt.Sprint(t)
}
func StINT32(t int32) string {
	return fmt.Sprint(t)
}
