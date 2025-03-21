package main

import "github.com/veandco/go-sdl2/sdl"

var ()

func UP() {

	INP()
	uBLOKS()
	uPL()
	uFX()
}
func uFX() {

}
func uBLOKS() {

	for i := range bloks {
		if bloks[i].velx != 0 || bloks[i].vely != 0 {
			if bloks[i].velx != 0 {
				if checkmoveblokX(i) {
					bloks[i].rd.X += bloks[i].velx
				} else {
					bloks[i].velx *= -1
				}
				if bloks[i].velx < 0 {
					if bloks[i].rd.X+bloks[i].rd.W < 0 {
						bloks[i].rd.X = DRAWW
					}
				}
				if bloks[i].velx > 0 {
					if bloks[i].rd.X > DRAWW {
						bloks[i].rd.X = -bloks[i].rd.W
					}
				}
			}
			if bloks[i].vely != 0 {
				if checkmoveblokY(i) {
					bloks[i].rd.Y += bloks[i].vely
				} else {
					bloks[i].vely *= -1
				}
				if bloks[i].vely < 0 {
					if bloks[i].rd.Y <= 0 {
						bloks[i].vely *= -1
					}
				}
				if bloks[i].vely > 0 {
					if bloks[i].rd.Y+bloks[i].rd.W >= yf {
						bloks[i].vely *= -1
					}
				}
			}
		}
	}

}
func checkmoveblokY(i int) bool {
	canmove := true

	b := bloks[i].rd
	b.Y += bloks[i].vely

	for j := range bloks {
		if i != j {
			if RECSINTERSECT(b, bloks[j].rd) && bloks[j].vely == 0 {
				canmove = false
			}
		}
	}

	return canmove
}
func checkmoveblokX(i int) bool {
	canmove := true

	b := bloks[i].rd
	b.X += bloks[i].velx

	for j := range bloks {
		if i != j {
			if RECSINTERSECT(b, bloks[j].rd) && bloks[j].velx == 0 {
				canmove = false
			}
		}
	}

	return canmove
}
func uPL() {

	//INP
	if kU && !pl.jump && pl.ground { //JUMP
		pl.state = 2
		pl.jump = true
		pl.jumpy1 = pl.cnt.Y
		pl.jumpy2 = pl.jumpy1 - pl.jumph
	}
	if pl.jump {
		pl.state = 2
		if pl.cnt.Y > pl.jumpy2 {
			if pl.vely > -pl.spd {
				pl.vely -= pl.acc * 2
			}
		} else if pl.cnt.Y <= pl.jumpy2 {
			pl.vely = 0
			pl.jump = false

		}
	}

	if kR { //LR
		pl.stop = false
		if pl.velx < pl.spd {
			pl.velx += 0.5
		}
	} else if kL {
		pl.stop = false
		if pl.velx > -pl.spd {
			pl.velx -= 0.5
		}
	}

	//LR
	if pl.velx > 0 {
		pl.lr = false
	} else if pl.velx < 0 {
		pl.lr = true
	}

	//STOP
	if kD {
		pl.stop = true
	}
	if pl.stop {
		if pl.velx > 0 {
			pl.velx -= pl.acc
			if pl.velx < 0 {
				pl.velx = 0
			}
		} else if pl.velx < 0 {
			pl.velx += pl.acc
			if pl.velx > 0 {
				pl.velx = 0
			}
		}
	}

	//STATE
	if pl.velx == 0 && pl.vely == 0 && !pl.jump {
		pl.state = 0
	} else if pl.velx != 0 || pl.vely != 0 && !pl.jump {
		pl.state = 1
	} else if pl.jump {
		pl.state = 2
	}

	//UP REC
	pl.cnt.X += float64(pl.velx)
	if pl.cnt.Y > -float64(pl.rc.H) {
		pl.cnt.Y += float64(pl.vely)
	}
	if pl.cnt.Y < 0 {
		pl.jump = false
		pl.ground = false
		pl.vely += pl.acc * 2

	}
	pl.r = sdl.Rect{int32(pl.cnt.X) - pl.r.W/2, int32(pl.cnt.Y) - pl.r.W/2, pl.r.W, pl.r.W}
	pl.rc = pl.r
	pl.rc.W -= (pl.r.W / 8) * 5
	pl.rc.X += pl.r.W / 4
	pl.rc.H -= pl.r.H / 5
	pl.rc.Y += pl.rc.H / 5

	//FLOOR GRAV
	if !pl.jump {
		if checkplblokfall() {
			if pl.cnt.Y+float64(pl.rc.H/2) < float64(yf) {
				if pl.vely < pl.spd {
					pl.vely += pl.acc * 2
				}
				if pl.cnt.Y+float64(pl.rc.H/2) >= float64(yf) {
					pl.cnt.Y = float64(yf) - float64(pl.rc.H/2)
					pl.vely = 0
				}
			} else {
				pl.vely = 0
			}
		}

	}

	if pl.cnt.Y+float64(pl.rc.H/2) < float64(yf) && !pl.onblock {
		pl.state = 2
		pl.ground = false
	} else {
		pl.ground = true
	}

	//LOOP SCREEN
	if pl.cnt.X-float64(pl.r.W/2) > float64(DRAWW) {
		pl.cnt.X = -float64(pl.r.W / 2)
	}
	if pl.cnt.X+float64(pl.r.W/2) < 0 {
		pl.cnt.X = float64(DRAWW) + float64(pl.r.W/2)
	}

}
func checkplblokfall() bool {

	falling := true

	r := pl.rc
	r.Y += 1
	for i := range bloks {
		if RECSINTERSECT(r, bloks[i].rd) && bloks[i].rd.Y > pl.rc.Y {
			falling = false
			pl.jump = false
			pl.onblock = true
			pl.vely = 0
			pl.cnt.Y = float64(bloks[i].rd.Y - pl.rc.H/2)
			if bloks[i].velx != 0 {
				if !kR && !kL {
					pl.state = 0
					pl.velx = float32(bloks[i].velx)
				}
			}
			break
		}
	}

	return falling
}
