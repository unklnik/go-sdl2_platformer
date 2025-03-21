package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	FT, WT, GT, BG, GUNIM, ETC []IM

	LOADTEX  *sdl.Texture
	LOADSURF *sdl.Surface
)

type IM struct {
	r, rd      sdl.Rect
	tex        *sdl.Texture
	velx, vely int32
}
type ANIM struct {
	r       sdl.Rect
	tex     *sdl.Texture
	imSLICE []IM

	frames, frameCOUNT, frameSTART, frameEND, frameCURR, xl int32
}

func mIM() {

	LOADSURF, _ = img.Load("im/etc.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)

	FT = mIMSLICE(0, 0, 16, 31, LOADTEX)
	WT = mIMSLICE(0, 16, 16, 39, LOADTEX)
	GT = mIMSLICE(0, 32, 64, 4, LOADTEX)
	D1IM = mIMSLICE(0, 96, 24, 24, LOADTEX)
	GUNIM = mIMSLICE(512, 0, 16, 7, LOADTEX)

	im := IM{}
	im.r = sdl.Rect{624, 0, 16, 16}
	im.tex = LOADTEX
	ETC = append(ETC, im) //0 TARGET

	LOADSURF, _ = img.Load("im/bg1.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg2.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg3.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg4.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg5.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg6.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)
	LOADSURF, _ = img.Load("im/bg7.png")
	LOADTEX, _ = RND.CreateTextureFromSurface(LOADSURF)
	im = IM{}
	im.r = sdl.Rect{0, 0, 1600, 900}
	im.tex = LOADTEX
	BG = append(BG, im)

}

func mANIMIMSLICE(im []IM, frameSTART, frameEND int32) ANIM {
	a := ANIM{}
	a.frameCURR = frameSTART
	a.frameSTART = frameSTART
	a.frameEND = frameEND
	a.imSLICE = im
	return a
}

func mIMSLICE(x, y, siz, num int32, tex *sdl.Texture) []IM {
	var im []IM
	for range num {
		im2 := IM{}
		im2.tex = tex
		im2.r = sdl.Rect{x, y, siz, siz}
		im = append(im, im2)
		x += siz
	}
	return im
}
