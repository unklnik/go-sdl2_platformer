package main

import "github.com/veandco/go-sdl2/sdl"

var ()

func dANIMIMSLICE(a ANIM, r sdl.Rect, spd int) ANIM {
	RND.Copy(a.imSLICE[0].tex, &a.imSLICE[a.frameCURR].r, &r)
	if FRAMES%spd == 0 {
		a.frameCURR++
		if a.frameCURR == a.frameEND {
			a.frameCURR = a.frameSTART
		}
	}
	return a
}
func dANIMIMSLICEFLIP(a ANIM, r sdl.Rect, spd int) ANIM {
	RND.CopyEx(a.imSLICE[0].tex, &a.imSLICE[a.frameCURR].r, &r, 0, ROCNTR(r), sdl.FLIP_HORIZONTAL)
	if FRAMES%spd == 0 {
		a.frameCURR++
		if a.frameCURR == a.frameEND {
			a.frameCURR = a.frameSTART
		}
	}
	return a
}
func dGRID(c sdl.Color, a uint8) {
	c.A = a
	for i := range grid {
		dRLINE(grid[i], c, 50, 1)
	}
}
func dLINEPOINTS(p []sdl.Point, c sdl.Color, a uint8) {
	for i := 0; i < len(p)-1; i += 2 {
		CA(c, a)
		rLINE(p[i], p[i+1])
	}
}
func dIM(i IM, r sdl.Rect) { //IMG
	RND.Copy(i.tex, &i.r, &r)
}
func dIMRO(i IM, r sdl.Rect, angl float64) {
	RND.CopyEx(i.tex, &i.r, &r, angl, ROCNTR(r), sdl.FLIP_NONE)
}
func dIMROFLIP(i IM, r sdl.Rect, angl float64) {
	RND.CopyEx(i.tex, &i.r, &r, angl, ROCNTR(r), sdl.FLIP_HORIZONTAL)
}
func dIMSLICE(im []IM, x, y, siz int32) {
	for i := range im {
		RND.Copy(im[i].tex, &im[i].r, &sdl.Rect{x, y, siz, siz})
		x += siz
	}
}
func dIMSLICERECOUTLINE(im []IM, c sdl.Color, a uint8, w int) {
	for i := range im {
		dRLINE(im[i].rd, c, a, w)
	}
}
func dIMSLICEREC(im []IM) {
	for i := range im {
		RND.Copy(im[i].tex, &im[i].r, &im[i].rd)
	}
}
func dIMSLICERECSHADOW(im []IM, offx, offy int32, a uint8) {
	for i := range im {
		r2 := im[i].rd
		r2.X -= offx
		r2.Y += offy
		c := DARKGREY()
		im[i].tex = CTEX(im[i].tex, c, a)
		RND.Copy(im[i].tex, &im[i].r, &r2)
		im[i].tex = TEXREVERT(im[i].tex)
		RND.Copy(im[i].tex, &im[i].r, &im[i].rd)
	}
}

func dRLINE(r sdl.Rect, c sdl.Color, a uint8, w int) { //REC
	CA(c, a)
	rRLINE(r, w)
}
func dR(r sdl.FRect, c sdl.Color) { //REC
	C(c)
	rR(r)
}
func dRA(r sdl.FRect, c sdl.Color, a uint8) { //REC ALPHA
	c.A = a
	C(c)
	rR(r)
}
func dRAI32(r sdl.Rect, c sdl.Color, a uint8) { //REC ALPHA
	c.A = a
	C(c)
	rRI32(r)
}
func rRLINE(r sdl.Rect, w int) { //RENDER REC LINE WIDTH
	RND.DrawRect(&r)
	if w > 1 {
		for w > 0 {
			r.X++
			r.Y++
			r.W -= 2
			r.H -= 2
			RND.DrawRect(&r)
			w--
		}

	}
}
func rLINE(p1, p2 sdl.Point) {
	RND.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}
func rR(r sdl.FRect) { //RENDER REC
	RND.FillRectF(&r)
}
func rRI32(r sdl.Rect) { //RENDER REC
	RND.FillRect(&r)
}
func C(c sdl.Color) { //SET DRAWCOLOR
	RND.SetDrawColor(c.R, c.G, c.B, c.A)
}
func CA(c sdl.Color, a uint8) { //SET DRAWCOLOR
	c.A = a
	RND.SetDrawColor(c.R, c.G, c.B, c.A)
}

func CTEX(t *sdl.Texture, c sdl.Color, a uint8) *sdl.Texture {
	t.SetColorMod(c.R, c.G, c.B)
	t.SetAlphaMod(a)
	return t
}
func TEXREVERT(t *sdl.Texture) *sdl.Texture {
	t.SetColorMod(255, 255, 255)
	t.SetAlphaMod(255)
	return t
}
