package main

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	FONSPC, FONLINEH float32 = 2, 2
	TXSURF           *sdl.Surface
	FON1PATH         = "fnt/Rubik-Medium.ttf"
	FON2PATH         = "fnt/Milonga-Regular.ttf"
	FON1, FON2       *ttf.Font

	FON1SMLR, FON1SML, FON1MED, FON1LRG, FON2SMLR, FON2SML, FON2MED, FON2LRG []CHAR

	tx14, tx18, tx24, tx30 = 14, 18, 24, 30

	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "
)

type CHAR struct {
	ch  string
	tex *sdl.Texture
	r   sdl.Rect
}

func dTXTxy(tx string, x, y float32, siz int, fon int) {
	t := strings.Split(tx, "")
	for i := range t {
		for j := range FON1MED {
			if t[i] == FON1MED[j].ch {
				switch fon {
				case 1: //FON1
					switch siz {
					case 1: //SMLR
						RND.CopyF(FON1SMLR[j].tex, &FON1SMLR[j].r, RǁPOINTER(x, y, float32(FON1SMLR[j].r.W), float32(FON1SMLR[j].r.H)))
						x += float32(FON1SMLR[j].r.W) + FONSPC
						break
					case 2: //SML
						RND.CopyF(FON1SML[j].tex, &FON1SML[j].r, RǁPOINTER(x, y, float32(FON1SML[j].r.W), float32(FON1SML[j].r.H)))
						x += float32(FON1SML[j].r.W) + FONSPC
						break
					case 3: //MED
						RND.CopyF(FON1MED[j].tex, &FON1MED[j].r, RǁPOINTER(x, y, float32(FON1MED[j].r.W), float32(FON1MED[j].r.H)))
						x += float32(FON1MED[j].r.W) + FONSPC
						break
					case 4: //LRG
						RND.CopyF(FON1LRG[j].tex, &FON1LRG[j].r, RǁPOINTER(x, y, float32(FON1LRG[j].r.W), float32(FON1LRG[j].r.H)))
						x += float32(FON1LRG[j].r.W) + FONSPC
						break
					}
				case 2: //FON2
					switch siz {
					case 1: //SMLR
						RND.CopyF(FON2SMLR[j].tex, &FON2SMLR[j].r, RǁPOINTER(x, y, float32(FON2SMLR[j].r.W), float32(FON2SMLR[j].r.H)))
						x += float32(FON2SMLR[j].r.W) + FONSPC
						break
					case 2: //SML
						RND.CopyF(FON2SML[j].tex, &FON2SML[j].r, RǁPOINTER(x, y, float32(FON2SML[j].r.W), float32(FON2SML[j].r.H)))
						x += float32(FON2SML[j].r.W) + FONSPC
						break
					case 3: //MED
						RND.CopyF(FON2MED[j].tex, &FON2MED[j].r, RǁPOINTER(x, y, float32(FON2MED[j].r.W), float32(FON2MED[j].r.H)))
						x += float32(FON2MED[j].r.W) + FONSPC
						break
					case 4: //LRG
						RND.CopyF(FON2LRG[j].tex, &FON2LRG[j].r, RǁPOINTER(x, y, float32(FON2LRG[j].r.W), float32(FON2LRG[j].r.H)))
						x += float32(FON2LRG[j].r.W) + FONSPC
						break
					}
				}
			}
		}

		/*
			for j := 0; j < len(fontchars); j++ {
				if t[i] == fontchars[j].character {
					RNDR.Copy(fontchars[j].tex, &fontchars[j].rec, &sdl.Rect{x, y, fontchars[j].rec.W, fontchars[j].rec.H})
					x += fontchars[j].rec.W + LetterSpace
					break
				}
			}
		*/
	}
}

func mFONTS() {
	ttf.Init()
	t := strings.Split(standardCharacters, "")
	FON1, _ = ttf.OpenFont(FON1PATH, tx14)
	for i := range t {
		FON1SMLR = append(FON1SMLR, mCHAR(t[i], FON1))
	}
	FON1, _ = ttf.OpenFont(FON1PATH, tx18)
	for i := range t {
		FON1SML = append(FON1SML, mCHAR(t[i], FON1))
	}
	FON1, _ = ttf.OpenFont(FON1PATH, tx24)
	for i := range t {
		FON1MED = append(FON1MED, mCHAR(t[i], FON1))
	}
	FON1, _ = ttf.OpenFont(FON1PATH, tx30)
	for i := range t {
		FON1LRG = append(FON1LRG, mCHAR(t[i], FON1))
	}
	FON2, _ = ttf.OpenFont(FON2PATH, tx14)
	for i := range t {
		FON2SMLR = append(FON2SMLR, mCHAR(t[i], FON2))
	}
	FON2, _ = ttf.OpenFont(FON2PATH, tx18)
	for i := range t {
		FON2SML = append(FON2SML, mCHAR(t[i], FON2))
	}
	FON2, _ = ttf.OpenFont(FON2PATH, tx24)
	for i := range t {
		FON2MED = append(FON2MED, mCHAR(t[i], FON2))
	}
	FON2, _ = ttf.OpenFont(FON2PATH, tx30)
	for i := range t {
		FON2LRG = append(FON2LRG, mCHAR(t[i], FON2))
	}

}

func mCHAR(c string, f *ttf.Font) CHAR {
	var w, h int
	TXSURF, _ = f.RenderUTF8Blended(c, WHITE())
	defer TXSURF.Free()
	w, h, _ = f.SizeUTF8(c)
	t := CHAR{}
	t.tex, _ = RND.CreateTextureFromSurface(TXSURF)
	t.r = sdl.Rect{0, 0, int32(w), int32(h)}
	t.ch = c
	return t
}
