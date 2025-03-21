package main

import "github.com/veandco/go-sdl2/sdl"

var ()

func R(x, y, w, h float32) sdl.FRect {
	return sdl.FRect{x, y, w, h}
}
func RǁPOINTER(x, y, w, h float32) *sdl.FRect {
	return &sdl.FRect{x, y, w, h}
}
func RCNT(w, h float32) sdl.FRect {
	return sdl.FRect{CNT.X - w/2, CNT.Y - h/2, w, h}
}
func SQ(x, y, w float32) sdl.FRect {
	return sdl.FRect{x, y, w, w}
}
func SQCNT(w float32) sdl.FRect {
	return sdl.FRect{CNT.X - w/2, CNT.Y - w/2, w, w}
}
