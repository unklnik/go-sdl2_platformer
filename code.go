package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	floorim, wallim, bgim, moveplatim IM
	grassIM, D1IM, bloks              []IM

	guns []GUN

	grid []sdl.Rect
	yf   int32
	pl   PLYR
)

type PLYR struct {
	r, rc     sdl.Rect
	cnt, targ Vector2
	state     int

	lr, stop, jump, ground, onblock bool

	idl, run, jumpanm ANIM

	velx, vely, spd, acc float32

	jumpy1, jumpy2, jumph, radtarg, gunangl, targangl float64

	gun GUN
}

type GUN struct {
	nm         string
	dmg, imnum int
	vel        float32
}

func DRAW() {

	dIM(bgim, DRAWREC)
	dRAI32(DRAWREC, BLACK(), 200)
	//FLOOR
	for i := len(grid) - TW*2; i < len(grid); i++ {
		dIM(floorim, grid[i])
	}
	dIMSLICE(grassIM, 0, yf-TSIZ, TSIZ*2)
	dIMSLICEREC(bloks)
	dIMSLICERECOUTLINE(bloks, BLACK(), 255, 2)

	//PLAYER
	dPL()

	//DEBUG
	if dbg {
		dGRID(GREEN(), 50)
		dRLINE(pl.r, MAGENTA(), 100, 2)
		dRLINE(pl.rc, ORANGE(), 100, 2)
	}

}
func dPL() { //MARK:dPL

	switch pl.state {
	case 2: //JUMP
		if pl.lr {
			pl.jumpanm = dANIMIMSLICEFLIP(pl.jumpanm, pl.r, 5)
		} else {
			pl.jumpanm = dANIMIMSLICE(pl.jumpanm, pl.r, 5)
		}
	case 1: //RUN
		if pl.lr {
			pl.run = dANIMIMSLICEFLIP(pl.run, pl.r, 4)
		} else {
			pl.run = dANIMIMSLICE(pl.run, pl.r, 4)
		}
	case 0: //IDLE
		if pl.lr {
			pl.idl = dANIMIMSLICEFLIP(pl.idl, pl.r, 8)
		} else {
			pl.idl = dANIMIMSLICE(pl.idl, pl.r, 8)
		}
	}

	//GUN
	gunrec := pl.r
	gunrec.W -= pl.r.W / 2
	gunrec.H -= pl.r.W / 2
	gunrec.X += pl.r.W / 4
	gunrec.Y += pl.r.W / 4
	gunrec.Y += gunrec.H / 5

	pl.gunangl = ANGL2POINTS(RECCENTERV2(gunrec), pl.targ)

	if pl.lr {
		pl.gunangl += -180
		gunrec.X -= (gunrec.W / 5) * 3
		dIMROFLIP(GUNIM[pl.gun.imnum], gunrec, pl.gunangl)
	} else {
		if pl.gunangl < 270 && pl.gunangl > 90 {
			pl.gunangl = 270
		} else if pl.gunangl < -90 {
			pl.gunangl = -90
		}
		gunrec.X += (gunrec.W / 5) * 3
		dIMRO(GUNIM[pl.gun.imnum], gunrec, pl.gunangl)
	}
	if dbg {
		dRLINE(gunrec, GREEN(), 255, 2)
	}

	//TARG
	pl.targangl = ANGL2POINTS(pl.cnt, MS)
	if pl.lr {

	} else {

	}
	pl.targ = POINTONCIRC(pl.radtarg, pl.targangl, pl.cnt)
	siz := TSIZ
	r := sdl.Rect{int32(pl.targ.X) - siz/2, int32(pl.targ.Y) - siz/2, siz, siz}
	dIM(ETC[0], r)

}
func mGUNS() {

	g := GUN{}

	for i := range 7 {

		switch i {
		case 0: //PISTOL
			g.nm = "Piztil"
			g.dmg = 1
			g.vel = float32(TSIZ / 4)
			g.imnum = 0

		}

		guns = append(guns, g)

	}

}

func mPL() { //MARK: mPL
	pl = PLYR{}
	siz := TSIZ * 2

	pl.idl = mANIMIMSLICE(D1IM, 0, 4)
	pl.run = mANIMIMSLICE(D1IM, 5, 11)
	pl.jumpanm = mANIMIMSLICE(D1IM, 12, 13)

	pl.cnt.X = float64(CNT.X)
	pl.cnt.Y = float64(yf - siz/2)
	pl.r = sdl.Rect{int32(pl.cnt.X) - siz/2, int32(pl.cnt.Y) - siz/2, siz, siz}
	pl.rc = pl.r
	pl.rc.W -= (pl.r.W / 8) * 5
	pl.rc.X += pl.r.W / 4
	pl.rc.H -= pl.r.H / 5
	pl.rc.Y += pl.rc.H / 5

	pl.spd = float32(TSIZ / 5)
	pl.acc = 0.5
	pl.jumph = float64(TSIZ) * 4

	pl.gun = guns[0]

	pl.radtarg = float64(TSIZ * 4)

}
func checkaddblok(r sdl.Rect) bool {
	canadd := true
	if len(bloks) > 0 {
		for i := range bloks {
			if RECSINTERSECT(r, bloks[i].rd) {
				canadd = false
				break
			}
		}
	}
	return canadd
}
func mLEV() { //MARK: mLEV

	bgim = BG[RINT(0, len(BG))]

	//PLATFORMS
	wallim = WT[RINT(0, len(WT))]

	var blokrange []int
	for i := TW; i < len(grid)-TW*4; i++ {
		blokrange = append(blokrange, i)
	}

	numplats := RINT(4, 8)

	for numplats > 0 {
		choose := RINT(0, 13)
		//choose = 12

		im := IM{}
		im.r = wallim.r
		im.tex = wallim.tex

		switch choose {
		case 12: //CROSS LINE
			num := RINT(3, 6)
			b := len(blokrange)
			b += RINT(0, TW-(num*4))
			b -= RINT(1, 5) * TW
			for num > 0 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b--
				b -= TW
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b += TW * 2
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b -= TW
				b += 3
				num--
			}
		case 11: //TRIANGLE DOWN
			siz := RINT(5, 10)
			if siz%2 == 0 {
				siz++
			}
			b := len(blokrange)
			b += RINT(0, TW-siz)
			b -= TW * siz / 2
			b -= RINT(0, 4) * TW
			ob := b

			for siz > 0 {
				for range siz {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				b = ob
				b += TW
				b++
				ob = b
				siz -= 2
			}

		case 10: //CROSS
			b := len(blokrange)
			b += RINT(0, TW-6)
			b -= 6 * TW
			ob := b

			for range 6 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b = ob
				b += TW
				ob = b
			}
			b -= 4 * TW
			b -= RINT(0, 5)
			ob = b
			for range 6 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
			}
			b = ob
			b += TW
			for range 6 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
			}

		case 9: //SOLID BLOCK
			siz := RINT(3, 8)
			b := len(blokrange)
			b += RINT(0, TW-siz)
			b -= siz * TW
			ob := b
			a := siz * siz
			c := 0
			for a > 0 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				c++
				if c == siz {
					c = 0
					b = ob
					b += TW
					ob = b
				}
				a--
			}

		case 8: //1 STEP LEFT
			num := RINT(3, 9)
			b := len(blokrange)
			b += RINT(num, TW)
			for num > 0 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b--
				b -= TW
				if b < 0 {
					break
				}
				num--
			}
		case 7: //1 STEP RIGHT
			num := RINT(3, 9)
			b := len(blokrange)
			b += RINT(0, TW/2)
			for num > 0 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				b -= TW
				if b < 0 {
					break
				}
				num--
			}

		case 6: //2 BLOK 1 SPACE
			num := RINT(2, 9)
			b := len(blokrange)
			b += RINT(0, TW/2)
			b -= RINT(0, TW*TH/2)

			for num > 0 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b++
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
				b += 2
				num--
			}

		case 5: //SMALL PYRAMID INVERT LEFT
			num := RINT(2, 9)
			leng := RINT(5, 11)
			leng2 := 2

			b := len(blokrange)
			b += RINT(0, TW-leng)
			ob := b

			for num > 0 && b > blokrange[0] {
				for range leng2 {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b--
				}
				if leng2 < leng {
					leng2++
				}
				b = ob
				b -= TW * 3
				ob = b
				num--
			}
		case 4: //SMALL PYRAMID INVERT RIGHT
			num := RINT(2, 9)
			leng := RINT(5, 11)
			leng2 := 2

			b := len(blokrange)
			b += RINT(0, TW-leng)
			ob := b

			for num > 0 && b > blokrange[0] {
				for range leng2 {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				if leng2 < leng {
					leng2++
				}
				b = ob
				b -= TW * 3
				ob = b
				num--
			}
		case 3: //SMALL PYRAMID RIGHT
			num := RINT(2, 9)
			leng := RINT(5, 11)

			b := len(blokrange)
			b += RINT(0, TW-leng)
			ob := b

			for num > 0 && b > blokrange[0] {
				for range leng {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				leng--
				b = ob
				b -= TW * 3
				b++
				ob = b
				num--
			}
		case 2: //SMALL PYRAMID LEFT
			num := RINT(2, 9)
			leng := RINT(5, 11)

			b := len(blokrange)
			b += RINT(0, TW-leng)
			ob := b

			for num > 0 && b > blokrange[0] {
				for range leng {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				leng--
				b = ob
				b -= TW * 3
				ob = b
				num--
			}
		case 1: //SMALL UNIFORM
			num := RINT(2, 9)
			leng := RINT(1, 8)

			b := len(blokrange)
			b += RINT(0, TW-leng)
			ob := b

			for num > 0 && b > blokrange[0] {
				for range leng {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				b = ob
				b -= TW * 3
				ob = b
				num--
			}

		case 0: //TRIANGLE UP
			siz := RINT(5, 10)
			if siz%2 == 0 {
				siz++
			}
			b := len(blokrange)
			b += RINT(0, TW-siz)
			b -= RINT(0, 4) * TW
			ob := b

			for siz > 0 {
				for range siz {
					im.rd = grid[b]
					if checkaddblok(im.rd) {
						bloks = append(bloks, im)
					}
					b++
				}
				b = ob
				b -= TW
				b++
				ob = b
				siz -= 2
			}

		}

		numplats--
	}
	//TOP LEVEL

	choose := RINT(0, 3)
	im := IM{}
	im.r = wallim.r
	im.tex = wallim.tex

	switch choose {
	case 2: //TRIANGLE
		b := TW * 3
		c := TW
		for c > 0 {
			if b < TW*4 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			c--
			b++
			b -= TW
			if b < TW*4 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			c--
			b++
			b += TW
			if b < TW*4 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b += 2
			c -= 2
		}
	case 1: //4 BLOK 1 SPACE
		b := TW * 2
		c := TW
		for c > 0 {
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b++
			c--
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b++
			c--
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b++
			c--
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b += 2
			c -= 2
		}
	case 0: //2 BLOK 1 SPACE
		b := TW * 2
		c := TW
		for c > 0 {
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b++
			c--
			if b < TW*3 {
				im.rd = grid[b]
				if checkaddblok(im.rd) {
					bloks = append(bloks, im)
				}
			}
			b += 2
			c -= 2
		}

	}

	//MOVING
	//LR
	for {
		moveplatim = WT[RINT(0, len(WT))]
		if moveplatim != wallim {
			break
		}
	}
	im = IM{}
	im.velx = RI32(2, 6)
	im.r = moveplatim.r
	im.tex = wallim.tex
	b := TW * RINT(4, 10)
	b += RINT(0, TW/2)
	siz := RINT(2, 8)
	for siz > 0 {
		im.rd = grid[b]
		if checkaddblok(im.rd) {
			bloks = append(bloks, im)
		}
		b++
		siz--
	}
	if FLIPCOIN() {
		im = IM{}
		im.velx = -RI32(2, 6)
		im.r = moveplatim.r
		im.tex = wallim.tex
		b := TW * RINT(10, 20)
		siz := RINT(2, 8)
		b += RINT(TW/2, TW-siz)
		for siz > 0 {
			im.rd = grid[b]
			if checkaddblok(im.rd) {
				bloks = append(bloks, im)
			}
			b++
			siz--
		}
	}
	//UD
	num := RINT(2, 8)
	for num > 0 {
		im = IM{}
		im.vely = RI32(2, 6)
		im.r = moveplatim.r
		im.tex = wallim.tex
		b = blokrange[RINT(0, len(blokrange))]
		im.rd = grid[b]
		if checkaddblok(im.rd) {
			bloks = append(bloks, im)
		}
		if FLIPCOIN() {
			b++
			im.rd = grid[b]
			if checkaddblok(im.rd) {
				bloks = append(bloks, im)
			}
		}
		num--
	}

	//FLOOR
	floorim = FT[RINT(0, len(FT))]
	yf = grid[len(grid)-TW*2].Y
	x := int32(0)
	for x < DRAWW {
		grassIM = append(grassIM, GT[RINT(0, len(GT))])
		x += TSIZ * 2
	}
}
