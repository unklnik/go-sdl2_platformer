package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (

	//KEYS
	kESC, kF1, kF2, kF3, kL, kR, kU, kD bool

	//MOUSE
	MS Vector2

	MSL, MSR, MSM bool

	MSCLL, MSCLR, MSCLM int
)

func INP() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch k := event.(type) {
		case sdl.QuitEvent:
			EXIT()

		case sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				switch k.Keysym.Scancode {
				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					kU = true
				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					kD = true
				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					kR = true
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					kL = true
				case sdl.SCANCODE_ESCAPE:
					kESC = true
				case sdl.SCANCODE_F1:
					kF1 = true
				case sdl.SCANCODE_F2:
					kF2 = true
				case sdl.SCANCODE_F3:
					kF3 = true
				}
			}
			if event.GetType() == sdl.KEYUP {
				switch k.Keysym.Scancode {
				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					kU = false
				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					kD = false
				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					kR = false
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					kL = false
				}
			}
		}
	}

	//KEYS
	if kESC {
		EXIT()
	}
	if kF3 {
		kF3 = false
	}
	if kF2 {
		RESTART()
		kF2 = false
	}
	if kF1 {
		dbg = !dbg
		kF1 = false
	}

	MOUSE()

}

func MOUSE() {
	x, y, click := sdl.GetMouseState()
	MS.X, MS.Y = float64(x), float64(y)

	if click == sdl.ButtonLMask && MSCLL == 0 {
		MSCLL = 4
		MSL = true
	} else {
		MSL = false
	}

	if click == sdl.ButtonRMask && MSCLR == 0 {
		MSCLR = 4
		MSR = true
	} else {
		MSR = false
	}

	if click == sdl.ButtonMMask && MSCLM == 0 {
		MSM = true
	} else {
		MSM = false
	}

}
