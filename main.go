package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	WIN             *sdl.Window
	RND             *sdl.Renderer
	SURF            *sdl.Surface
	TEX             *sdl.Texture
	DRAWREC, RNDREC sdl.Rect
	DISPLAYMODE     sdl.DisplayMode

	TSIZ   int32
	TDIV   = int32(24)
	TW, TH int

	FRAMES int

	CNT sdl.FPoint

	RNDW, RNDH, DRAWW, DRAWH int32 = 1920, 1080, 1920, 1080

	OFF, VSYNC bool
)

func RESTART() {
	grassIM = nil
	bloks = nil
	mLEV()

}

func PLAY() {

	sdl.ShowCursor(sdl.DISABLE)
	//SCANLINES = true

	mFONTS()
	mIM()
	mFX()
	mLEV()
	mGUNS()
	mPL()

	for !OFF {
		FRAMES++
		RND.SetRenderTarget(TEX)
		RND.SetDrawColor(0, 0, 0, 0)
		RND.Clear()

		DRAW()

		if dbg {
			DEBUG()
		}

		RND.SetRenderTarget(nil)
		RND.SetDrawColor(0, 0, 0, 0)
		RND.Clear()
		RND.Copy(TEX, &DRAWREC, &RNDREC)
		RND.Present()

		if !VSYNC {
			sdl.Delay(16)
		}

		UP()

	}

}
func EXIT() {
	LOADSURF.Free()
	LOADTEX.Destroy()
	TEX.Destroy()
	RND.Destroy()
	WIN.Destroy()
	OFF = true

}

func main() {
	VSYNC = true

	//GET SCREEN SIZE

	//CREATE WIN RND TEX
	WIN, _ = sdl.CreateWindow("SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, RNDW, RNDH, sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_SHOWN)
	RND, _ = sdl.CreateRenderer(WIN, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	RND.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	TEX, _ = RND.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, DRAWW, DRAWH)

	defer TEX.Destroy()
	defer RND.Destroy()
	defer WIN.Destroy()

	DRAWREC = sdl.Rect{0, 0, DRAWW, DRAWH}
	RNDREC = sdl.Rect{0, 0, RNDW, RNDH}
	TSIZ = DRAWH / TDIV
	TWfound := false
	var x, y int32 = 0, 0 //MAKE GRID
	for y < DRAWH {
		r := sdl.Rect{x, y, TSIZ, TSIZ}
		grid = append(grid, r)
		x += TSIZ
		if !TWfound {
			TW++
		}
		if x >= DRAWW {
			TWfound = true
			TH++
			x = 0
			y += TSIZ
		}
	}

	RND.SetRenderTarget(TEX)

	CNT = sdl.FPoint{float32(DRAWW / 2), float32(DRAWH / 2)}

	PLAY()
}
