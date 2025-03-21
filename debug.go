package main

import "fmt"

var (
	dbg bool
)

func DEBUG() {
	dRA(R(0, 0, 300, float32(RNDH)), MAROON(), 100)
	var x, y float32 = 4, 4
	dTXTxy("RNDW "+fmt.Sprint(RNDW)+" RNDH "+fmt.Sprint(RNDH), x, y, 1, 1)
	y += float32(FON1SMLR[0].r.H)
	dTXTxy("DRAWW "+fmt.Sprint(DRAWW)+" DRAWH "+fmt.Sprint(DRAWH), x, y, 1, 1)
	y += float32(FON1SMLR[0].r.H)
	dTXTxy("TW "+fmt.Sprint(TW)+" TH "+fmt.Sprint(TH)+" TSIZ "+fmt.Sprint(TSIZ), x, y, 1, 1)
	y += float32(FON1SMLR[0].r.H)
	dTXTxy("yf "+fmt.Sprint(yf)+" pl.gunangl "+fmt.Sprintf("%.0f", pl.gunangl)+" pl.targangl "+fmt.Sprintf("%.0f", pl.targangl), x, y, 1, 1)
	y += float32(FON1SMLR[0].r.H)
}
