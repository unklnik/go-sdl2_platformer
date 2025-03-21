package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// MARK: COLORS
func colorOpacity(c []uint8, opacity uint8) (uint8, uint8, uint8, uint8) {
	return c[0], c[1], c[2], opacity
}

func color2array(c1 uint8, c2 uint8, c3 uint8, c4 uint8) []uint8 {
	var colors []uint8
	colors = append(colors, c1, c2, c3, c4)
	return colors
}

// RANDOM COLORS
func COLORǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOM() sdl.Color {
	return sdl.Color{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

func COLORǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255
}
func DARKGREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255
}
func GREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255
}
func REDǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255
}
func PINKǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255
}
func BLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255
}
func DARKBLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255
}
func CYANǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255
}
func YELLOWǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255
}
func ORANGEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return 255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255
}
func BROWNǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255
}
func GREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255
}
func DARKGREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255
}

func COLORǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

// SOLID COLORS
func MAROON() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(0), 255} }
func DARKRED() sdl.Color              { return sdl.Color{uint8(139), uint8(0), uint8(0), 255} }
func BROWN() sdl.Color                { return sdl.Color{uint8(165), uint8(42), uint8(42), 255} }
func FIREBRICK() sdl.Color            { return sdl.Color{uint8(178), uint8(34), uint8(34), 255} }
func CRIMSON() sdl.Color              { return sdl.Color{uint8(220), uint8(20), uint8(60), 255} }
func RED() sdl.Color                  { return sdl.Color{uint8(255), uint8(0), uint8(0), 255} }
func TOMATO() sdl.Color               { return sdl.Color{uint8(255), uint8(99), uint8(71), 255} }
func CORAL() sdl.Color                { return sdl.Color{uint8(255), uint8(127), uint8(80), 255} }
func INDIANRED() sdl.Color            { return sdl.Color{uint8(205), uint8(92), uint8(92), 255} }
func LIGHTCORAL() sdl.Color           { return sdl.Color{uint8(240), uint8(128), uint8(128), 255} }
func DARKSALMON() sdl.Color           { return sdl.Color{uint8(233), uint8(150), uint8(122), 255} }
func SALMON() sdl.Color               { return sdl.Color{uint8(250), uint8(128), uint8(114), 255} }
func LIGHTSALMON() sdl.Color          { return sdl.Color{uint8(255), uint8(160), uint8(122), 255} }
func ORANGERED() sdl.Color            { return sdl.Color{uint8(255), uint8(69), uint8(0), 255} }
func DARKORANGE() sdl.Color           { return sdl.Color{uint8(255), uint8(140), uint8(0), 255} }
func ORANGE() sdl.Color               { return sdl.Color{uint8(255), uint8(165), uint8(0), 255} }
func GOLD() sdl.Color                 { return sdl.Color{uint8(255), uint8(215), uint8(0), 255} }
func DARKGOLDENROD() sdl.Color        { return sdl.Color{uint8(184), uint8(134), uint8(11), 255} }
func GOLDENROD() sdl.Color            { return sdl.Color{uint8(218), uint8(165), uint8(32), 255} }
func PALEGOLDENROD() sdl.Color        { return sdl.Color{uint8(238), uint8(232), uint8(170), 255} }
func DARKKHAKI() sdl.Color            { return sdl.Color{uint8(189), uint8(183), uint8(107), 255} }
func KHAKI() sdl.Color                { return sdl.Color{uint8(240), uint8(230), uint8(140), 255} }
func OLIVE() sdl.Color                { return sdl.Color{uint8(128), uint8(128), uint8(0), 255} }
func YELLOW() sdl.Color               { return sdl.Color{uint8(255), uint8(255), uint8(0), 255} }
func YELLOWGREEN() sdl.Color          { return sdl.Color{uint8(154), uint8(205), uint8(50), 255} }
func DARKOLIVEGREEN() sdl.Color       { return sdl.Color{uint8(85), uint8(107), uint8(47), 255} }
func OLIVEDRAB() sdl.Color            { return sdl.Color{uint8(107), uint8(142), uint8(35), 255} }
func LAWNGREEN() sdl.Color            { return sdl.Color{uint8(124), uint8(252), uint8(0), 255} }
func CHARTREUSE() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(0), 255} }
func GREENYELLOW() sdl.Color          { return sdl.Color{uint8(173), uint8(255), uint8(47), 255} }
func DARKGREEN() sdl.Color            { return sdl.Color{uint8(0), uint8(100), uint8(0), 255} }
func GREEN() sdl.Color                { return sdl.Color{uint8(0), uint8(128), uint8(0), 255} }
func FORESTGREEN() sdl.Color          { return sdl.Color{uint8(34), uint8(139), uint8(34), 255} }
func LIME() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(0), 255} }
func LIMEGREEN() sdl.Color            { return sdl.Color{uint8(50), uint8(205), uint8(50), 255} }
func LIGHTGREEN() sdl.Color           { return sdl.Color{uint8(144), uint8(238), uint8(144), 255} }
func PALEGREEN() sdl.Color            { return sdl.Color{uint8(152), uint8(251), uint8(152), 255} }
func DARKSEAGREEN() sdl.Color         { return sdl.Color{uint8(143), uint8(188), uint8(143), 255} }
func MEDIUMSPRINGGREEN() sdl.Color    { return sdl.Color{uint8(0), uint8(250), uint8(154), 255} }
func SPRINGGREEN() sdl.Color          { return sdl.Color{uint8(0), uint8(255), uint8(127), 255} }
func SEAGREEN() sdl.Color             { return sdl.Color{uint8(46), uint8(139), uint8(87), 255} }
func MEDIUMAQUAMARINE() sdl.Color     { return sdl.Color{uint8(102), uint8(205), uint8(170), 255} }
func MEDIUMSEAGREEN() sdl.Color       { return sdl.Color{uint8(60), uint8(179), uint8(113), 255} }
func LIGHTSEAGREEN() sdl.Color        { return sdl.Color{uint8(32), uint8(178), uint8(170), 255} }
func DARKSLATEGRAY() sdl.Color        { return sdl.Color{uint8(47), uint8(79), uint8(79), 255} }
func TEAL() sdl.Color                 { return sdl.Color{uint8(0), uint8(128), uint8(128), 255} }
func DARKCYAN() sdl.Color             { return sdl.Color{uint8(0), uint8(139), uint8(139), 255} }
func AQUA() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func CYAN() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func LIGHTCYAN() sdl.Color            { return sdl.Color{uint8(224), uint8(255), uint8(255), 255} }
func DARKTURQUOISE() sdl.Color        { return sdl.Color{uint8(0), uint8(206), uint8(209), 255} }
func TURQUOISE() sdl.Color            { return sdl.Color{uint8(64), uint8(224), uint8(208), 255} }
func MEDIUMTURQUOISE() sdl.Color      { return sdl.Color{uint8(72), uint8(209), uint8(204), 255} }
func PALETURQUOISE() sdl.Color        { return sdl.Color{uint8(175), uint8(238), uint8(238), 255} }
func AQUAMARINE() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(212), 255} }
func POWDERBLUE() sdl.Color           { return sdl.Color{uint8(176), uint8(224), uint8(230), 255} }
func CADETBLUE() sdl.Color            { return sdl.Color{uint8(95), uint8(158), uint8(160), 255} }
func STEELBLUE() sdl.Color            { return sdl.Color{uint8(70), uint8(130), uint8(180), 255} }
func CORNFLOWERBLUE() sdl.Color       { return sdl.Color{uint8(100), uint8(149), uint8(237), 255} }
func DEEPSKYBLUE() sdl.Color          { return sdl.Color{uint8(0), uint8(191), uint8(255), 255} }
func DODGERBLUE() sdl.Color           { return sdl.Color{uint8(30), uint8(144), uint8(255), 255} }
func LIGHTBLUE() sdl.Color            { return sdl.Color{uint8(173), uint8(216), uint8(230), 255} }
func SKYBLUE() sdl.Color              { return sdl.Color{uint8(135), uint8(206), uint8(235), 255} }
func LIGHTSKYBLUE() sdl.Color         { return sdl.Color{uint8(135), uint8(206), uint8(250), 255} }
func MIDNIGHTBLUE() sdl.Color         { return sdl.Color{uint8(25), uint8(25), uint8(112), 255} }
func NAVY() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(128), 255} }
func DARKBLUE() sdl.Color             { return sdl.Color{uint8(0), uint8(0), uint8(139), 255} }
func MEDIUMBLUE() sdl.Color           { return sdl.Color{uint8(0), uint8(0), uint8(205), 255} }
func BLUE() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(255), 255} }
func ROYALBLUE() sdl.Color            { return sdl.Color{uint8(65), uint8(105), uint8(225), 255} }
func BLUEVIOLET() sdl.Color           { return sdl.Color{uint8(138), uint8(43), uint8(226), 255} }
func INDIGO() sdl.Color               { return sdl.Color{uint8(75), uint8(0), uint8(130), 255} }
func DARKSLATEBLUE() sdl.Color        { return sdl.Color{uint8(72), uint8(61), uint8(139), 255} }
func SLATEBLUE() sdl.Color            { return sdl.Color{uint8(106), uint8(90), uint8(205), 255} }
func MEDIUMSLATEBLUE() sdl.Color      { return sdl.Color{uint8(123), uint8(104), uint8(238), 255} }
func MEDIUMPURPLE() sdl.Color         { return sdl.Color{uint8(147), uint8(112), uint8(219), 255} }
func DARKMAGENTA() sdl.Color          { return sdl.Color{uint8(139), uint8(0), uint8(139), 255} }
func DARKVIOLET() sdl.Color           { return sdl.Color{uint8(148), uint8(0), uint8(211), 255} }
func DARKORCHID() sdl.Color           { return sdl.Color{uint8(153), uint8(50), uint8(204), 255} }
func MEDIUMORCHID() sdl.Color         { return sdl.Color{uint8(186), uint8(85), uint8(211), 255} }
func PURPLE() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(128), 255} }
func THISTLE() sdl.Color              { return sdl.Color{uint8(216), uint8(191), uint8(216), 255} }
func PLUM() sdl.Color                 { return sdl.Color{uint8(221), uint8(160), uint8(221), 255} }
func VIOLET() sdl.Color               { return sdl.Color{uint8(238), uint8(130), uint8(238), 255} }
func MAGENTA() sdl.Color              { return sdl.Color{uint8(255), uint8(0), uint8(255), 255} }
func ORCHID() sdl.Color               { return sdl.Color{uint8(218), uint8(112), uint8(214), 255} }
func MEDIUMVIOLETRED() sdl.Color      { return sdl.Color{uint8(199), uint8(21), uint8(133), 255} }
func PALEVIOLETRED() sdl.Color        { return sdl.Color{uint8(219), uint8(112), uint8(147), 255} }
func DEEPPINK() sdl.Color             { return sdl.Color{uint8(255), uint8(20), uint8(147), 255} }
func HOTPINK() sdl.Color              { return sdl.Color{uint8(255), uint8(105), uint8(180), 255} }
func LIGHTPINK() sdl.Color            { return sdl.Color{uint8(255), uint8(182), uint8(193), 255} }
func PINK() sdl.Color                 { return sdl.Color{uint8(255), uint8(192), uint8(203), 255} }
func ANTIQUEWHITE() sdl.Color         { return sdl.Color{uint8(250), uint8(235), uint8(215), 255} }
func BEIGE() sdl.Color                { return sdl.Color{uint8(245), uint8(245), uint8(220), 255} }
func BISQUE() sdl.Color               { return sdl.Color{uint8(255), uint8(228), uint8(196), 255} }
func BLANCHEDALMOND() sdl.Color       { return sdl.Color{uint8(255), uint8(235), uint8(205), 255} }
func WHEAT() sdl.Color                { return sdl.Color{uint8(245), uint8(222), uint8(179), 255} }
func CORNSILK() sdl.Color             { return sdl.Color{uint8(255), uint8(248), uint8(220), 255} }
func LEMONCHIFFON() sdl.Color         { return sdl.Color{uint8(255), uint8(250), uint8(205), 255} }
func LIGHTGOLDENRODYELLOW() sdl.Color { return sdl.Color{uint8(250), uint8(250), uint8(210), 255} }
func LIGHTYELLOW() sdl.Color          { return sdl.Color{uint8(255), uint8(255), uint8(224), 255} }
func SADDLEBROWN() sdl.Color          { return sdl.Color{uint8(139), uint8(69), uint8(19), 255} }
func SIENNA() sdl.Color               { return sdl.Color{uint8(160), uint8(82), uint8(45), 255} }
func CHOCOLATE() sdl.Color            { return sdl.Color{uint8(210), uint8(105), uint8(30), 255} }
func PERU() sdl.Color                 { return sdl.Color{uint8(205), uint8(133), uint8(63), 255} }
func SANDYBROWN() sdl.Color           { return sdl.Color{uint8(244), uint8(164), uint8(96), 255} }
func BURLYWOOD() sdl.Color            { return sdl.Color{uint8(222), uint8(184), uint8(135), 255} }
func TAN() sdl.Color                  { return sdl.Color{uint8(210), uint8(180), uint8(140), 255} }
func ROSYBROWN() sdl.Color            { return sdl.Color{uint8(188), uint8(143), uint8(143), 255} }
func MOCCASIN() sdl.Color             { return sdl.Color{uint8(255), uint8(228), uint8(181), 255} }
func NAVAJOWHITE() sdl.Color          { return sdl.Color{uint8(255), uint8(222), uint8(173), 255} }
func PEACHPUFF() sdl.Color            { return sdl.Color{uint8(255), uint8(218), uint8(185), 255} }
func MISTYROSE() sdl.Color            { return sdl.Color{uint8(255), uint8(228), uint8(225), 255} }
func LAVENDERBLUSH() sdl.Color        { return sdl.Color{uint8(255), uint8(240), uint8(245), 255} }
func LINEN() sdl.Color                { return sdl.Color{uint8(250), uint8(240), uint8(230), 255} }
func OLDLACE() sdl.Color              { return sdl.Color{uint8(253), uint8(245), uint8(230), 255} }
func PAPAYAWHIP() sdl.Color           { return sdl.Color{uint8(255), uint8(239), uint8(213), 255} }
func SEASHELL() sdl.Color             { return sdl.Color{uint8(255), uint8(245), uint8(238), 255} }
func MINTCREAM() sdl.Color            { return sdl.Color{uint8(245), uint8(255), uint8(250), 255} }
func SLATEGRAY() sdl.Color            { return sdl.Color{uint8(112), uint8(128), uint8(144), 255} }
func LIGHTSLATEGRAY() sdl.Color       { return sdl.Color{uint8(119), uint8(136), uint8(153), 255} }
func LIGHTSTEELBLUE() sdl.Color       { return sdl.Color{uint8(176), uint8(196), uint8(222), 255} }
func LAVENDER() sdl.Color             { return sdl.Color{uint8(230), uint8(230), uint8(250), 255} }
func FLORALWHITE() sdl.Color          { return sdl.Color{uint8(255), uint8(250), uint8(240), 255} }
func ALICEBLUE() sdl.Color            { return sdl.Color{uint8(240), uint8(248), uint8(255), 255} }
func GHOSTWHITE() sdl.Color           { return sdl.Color{uint8(248), uint8(248), uint8(255), 255} }
func HONEYDEW() sdl.Color             { return sdl.Color{uint8(240), uint8(255), uint8(240), 255} }
func IVORY() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(240), 255} }
func AZURE() sdl.Color                { return sdl.Color{uint8(240), uint8(255), uint8(255), 255} }
func SNOW() sdl.Color                 { return sdl.Color{uint8(255), uint8(250), uint8(250), 255} }
func BLACK() sdl.Color                { return sdl.Color{uint8(0), uint8(0), uint8(0), 255} }
func DIMGREY() sdl.Color              { return sdl.Color{uint8(105), uint8(105), uint8(105), 255} }
func GREY() sdl.Color                 { return sdl.Color{uint8(128), uint8(128), uint8(128), 255} }
func DARKGREY() sdl.Color             { return sdl.Color{uint8(169), uint8(169), uint8(169), 255} }
func SILVER() sdl.Color               { return sdl.Color{uint8(192), uint8(192), uint8(192), 255} }
func LIGHTGREY() sdl.Color            { return sdl.Color{uint8(211), uint8(211), uint8(211), 255} }
func GAINSBORO() sdl.Color            { return sdl.Color{uint8(220), uint8(220), uint8(220), 255} }
func WHITESMOKE() sdl.Color           { return sdl.Color{uint8(245), uint8(245), uint8(245), 255} }
func WHITE() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(255), 255} }

func MAROONǁ2() (uint8, uint8, uint8, uint8)               { return 128, 0, 0, 255 }
func DARKREDǁ2() (uint8, uint8, uint8, uint8)              { return 139, 0, 0, 255 }
func BROWNǁ2() (uint8, uint8, uint8, uint8)                { return 165, 42, 42, 255 }
func FIREBRICKǁ2() (uint8, uint8, uint8, uint8)            { return 178, 34, 34, 255 }
func CRIMSONǁ2() (uint8, uint8, uint8, uint8)              { return 220, 20, 60, 255 }
func REDǁ2() (uint8, uint8, uint8, uint8)                  { return 255, 0, 0, 255 }
func TOMATOǁ2() (uint8, uint8, uint8, uint8)               { return 255, 99, 71, 255 }
func CORALǁ2() (uint8, uint8, uint8, uint8)                { return 255, 127, 80, 255 }
func INDIANREDǁ2() (uint8, uint8, uint8, uint8)            { return 205, 92, 92, 255 }
func LIGHTCORALǁ2() (uint8, uint8, uint8, uint8)           { return 240, 128, 128, 255 }
func DARKSALMONǁ2() (uint8, uint8, uint8, uint8)           { return 233, 150, 122, 255 }
func SALMONǁ2() (uint8, uint8, uint8, uint8)               { return 250, 128, 114, 255 }
func LIGHTSALMONǁ2() (uint8, uint8, uint8, uint8)          { return 255, 160, 122, 255 }
func ORANGEREDǁ2() (uint8, uint8, uint8, uint8)            { return 255, 69, 0, 255 }
func DARKORANGEǁ2() (uint8, uint8, uint8, uint8)           { return 255, 140, 0, 255 }
func ORANGEǁ2() (uint8, uint8, uint8, uint8)               { return 255, 165, 0, 255 }
func GOLDǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 215, 0, 255 }
func DARKGOLDENRODǁ2() (uint8, uint8, uint8, uint8)        { return 184, 134, 11, 255 }
func GOLDENRODǁ2() (uint8, uint8, uint8, uint8)            { return 218, 165, 32, 255 }
func PALEGOLDENRODǁ2() (uint8, uint8, uint8, uint8)        { return 238, 232, 170, 255 }
func DARKKHAKIǁ2() (uint8, uint8, uint8, uint8)            { return 189, 183, 107, 255 }
func KHAKIǁ2() (uint8, uint8, uint8, uint8)                { return 240, 230, 140, 255 }
func OLIVEǁ2() (uint8, uint8, uint8, uint8)                { return 128, 128, 0, 255 }
func YELLOWǁ2() (uint8, uint8, uint8, uint8)               { return 255, 255, 0, 255 }
func YELLOWGREENǁ2() (uint8, uint8, uint8, uint8)          { return 154, 205, 50, 255 }
func DARKOLIVEGREENǁ2() (uint8, uint8, uint8, uint8)       { return 85, 107, 47, 255 }
func OLIVEDRABǁ2() (uint8, uint8, uint8, uint8)            { return 107, 142, 35, 255 }
func LAWNGREENǁ2() (uint8, uint8, uint8, uint8)            { return 124, 252, 0, 255 }
func CHARTREUSEǁ2() (uint8, uint8, uint8, uint8)           { return 127, 255, 0, 255 }
func GREENYELLOWǁ2() (uint8, uint8, uint8, uint8)          { return 173, 255, 47, 255 }
func DARKGREENǁ2() (uint8, uint8, uint8, uint8)            { return 0, 100, 0, 255 }
func GREENǁ2() (uint8, uint8, uint8, uint8)                { return 0, 128, 0, 255 }
func FORESTGREENǁ2() (uint8, uint8, uint8, uint8)          { return 34, 139, 34, 255 }
func LIMEǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 0, 255 }
func LIMEGREENǁ2() (uint8, uint8, uint8, uint8)            { return 50, 205, 50, 255 }
func LIGHTGREENǁ2() (uint8, uint8, uint8, uint8)           { return 144, 238, 144, 255 }
func PALEGREENǁ2() (uint8, uint8, uint8, uint8)            { return 152, 251, 152, 255 }
func DARKSEAGREENǁ2() (uint8, uint8, uint8, uint8)         { return 143, 188, 143, 255 }
func MEDIUMSPRINGGREENǁ2() (uint8, uint8, uint8, uint8)    { return 0, 250, 154, 255 }
func SPRINGGREENǁ2() (uint8, uint8, uint8, uint8)          { return 0, 255, 127, 255 }
func SEAGREENǁ2() (uint8, uint8, uint8, uint8)             { return 46, 139, 87, 255 }
func MEDIUMAQUAMARINEǁ2() (uint8, uint8, uint8, uint8)     { return 102, 205, 170, 255 }
func MEDIUMSEAGREENǁ2() (uint8, uint8, uint8, uint8)       { return 60, 179, 113, 255 }
func LIGHTSEAGREENǁ2() (uint8, uint8, uint8, uint8)        { return 32, 178, 170, 255 }
func DARKSLATEGRAYǁ2() (uint8, uint8, uint8, uint8)        { return 47, 79, 79, 255 }
func TEALǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 128, 128, 255 }
func DARKCYANǁ2() (uint8, uint8, uint8, uint8)             { return 0, 139, 139, 255 }
func AQUAǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 255, 255 }
func CYANǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 255, 255 }
func LIGHTCYANǁ2() (uint8, uint8, uint8, uint8)            { return 224, 255, 255, 255 }
func DARKTURQUOISEǁ2() (uint8, uint8, uint8, uint8)        { return 0, 206, 209, 255 }
func TURQUOISEǁ2() (uint8, uint8, uint8, uint8)            { return 64, 224, 208, 255 }
func MEDIUMTURQUOISEǁ2() (uint8, uint8, uint8, uint8)      { return 72, 209, 204, 255 }
func PALETURQUOISEǁ2() (uint8, uint8, uint8, uint8)        { return 175, 238, 238, 255 }
func AQUAMARINEǁ2() (uint8, uint8, uint8, uint8)           { return 127, 255, 212, 255 }
func POWDERBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 176, 224, 230, 255 }
func CADETBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 95, 158, 160, 255 }
func STEELBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 70, 130, 180, 255 }
func CORNFLOWERBLUEǁ2() (uint8, uint8, uint8, uint8)       { return 100, 149, 237, 255 }
func DEEPSKYBLUEǁ2() (uint8, uint8, uint8, uint8)          { return 0, 191, 255, 255 }
func DODGERBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 30, 144, 255, 255 }
func LIGHTBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 173, 216, 230, 255 }
func SKYBLUEǁ2() (uint8, uint8, uint8, uint8)              { return 135, 206, 235, 255 }
func LIGHTSKYBLUEǁ2() (uint8, uint8, uint8, uint8)         { return 135, 206, 250, 255 }
func MIDNIGHTBLUEǁ2() (uint8, uint8, uint8, uint8)         { return 25, 25, 112, 255 }
func NAVYǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 0, 128, 255 }
func DARKBLUEǁ2() (uint8, uint8, uint8, uint8)             { return 0, 0, 139, 255 }
func MEDIUMBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 0, 0, 205, 255 }
func BLUEǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 0, 255, 255 }
func ROYALBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 65, 105, 225, 255 }
func BLUEVIOLETǁ2() (uint8, uint8, uint8, uint8)           { return 138, 43, 226, 255 }
func INDIGOǁ2() (uint8, uint8, uint8, uint8)               { return 75, 0, 130, 255 }
func DARKSLATEBLUEǁ2() (uint8, uint8, uint8, uint8)        { return 72, 61, 139, 255 }
func SLATEBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 106, 90, 205, 255 }
func MEDIUMSLATEBLUEǁ2() (uint8, uint8, uint8, uint8)      { return 123, 104, 238, 255 }
func MEDIUMPURPLEǁ2() (uint8, uint8, uint8, uint8)         { return 147, 112, 219, 255 }
func DARKMAGENTAǁ2() (uint8, uint8, uint8, uint8)          { return 139, 0, 139, 255 }
func DARKVIOLETǁ2() (uint8, uint8, uint8, uint8)           { return 148, 0, 211, 255 }
func DARKORCHIDǁ2() (uint8, uint8, uint8, uint8)           { return 153, 50, 204, 255 }
func MEDIUMORCHIDǁ2() (uint8, uint8, uint8, uint8)         { return 186, 85, 211, 255 }
func PURPLEǁ2() (uint8, uint8, uint8, uint8)               { return 128, 0, 128, 255 }
func THISTLEǁ2() (uint8, uint8, uint8, uint8)              { return 216, 191, 216, 255 }
func PLUMǁ2() (uint8, uint8, uint8, uint8)                 { return 221, 160, 221, 255 }
func VIOLETǁ2() (uint8, uint8, uint8, uint8)               { return 238, 130, 238, 255 }
func MAGENTAǁ2() (uint8, uint8, uint8, uint8)              { return 255, 0, 255, 255 }
func ORCHIDǁ2() (uint8, uint8, uint8, uint8)               { return 218, 112, 214, 255 }
func MEDIUMVIOLETREDǁ2() (uint8, uint8, uint8, uint8)      { return 199, 21, 133, 255 }
func PALEVIOLETREDǁ2() (uint8, uint8, uint8, uint8)        { return 219, 112, 147, 255 }
func DEEPPINKǁ2() (uint8, uint8, uint8, uint8)             { return 255, 20, 147, 255 }
func HOTPINKǁ2() (uint8, uint8, uint8, uint8)              { return 255, 105, 180, 255 }
func LIGHTPINKǁ2() (uint8, uint8, uint8, uint8)            { return 255, 182, 193, 255 }
func PINKǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 192, 203, 255 }
func ANTIQUEWHITEǁ2() (uint8, uint8, uint8, uint8)         { return 250, 235, 215, 255 }
func BEIGEǁ2() (uint8, uint8, uint8, uint8)                { return 245, 245, 220, 255 }
func BISQUEǁ2() (uint8, uint8, uint8, uint8)               { return 255, 228, 196, 255 }
func BLANCHEDALMONDǁ2() (uint8, uint8, uint8, uint8)       { return 255, 235, 205, 255 }
func WHEATǁ2() (uint8, uint8, uint8, uint8)                { return 245, 222, 179, 255 }
func CORNSILKǁ2() (uint8, uint8, uint8, uint8)             { return 255, 248, 220, 255 }
func LEMONCHIFFONǁ2() (uint8, uint8, uint8, uint8)         { return 255, 250, 205, 255 }
func LIGHTGOLDENRODYELLOWǁ2() (uint8, uint8, uint8, uint8) { return 250, 250, 210, 255 }
func LIGHTYELLOWǁ2() (uint8, uint8, uint8, uint8)          { return 255, 255, 224, 255 }
func SADDLEBROWNǁ2() (uint8, uint8, uint8, uint8)          { return 139, 69, 19, 255 }
func SIENNAǁ2() (uint8, uint8, uint8, uint8)               { return 160, 82, 45, 255 }
func CHOCOLATEǁ2() (uint8, uint8, uint8, uint8)            { return 210, 105, 30, 255 }
func PERUǁ2() (uint8, uint8, uint8, uint8)                 { return 205, 133, 63, 255 }
func SANDYBROWNǁ2() (uint8, uint8, uint8, uint8)           { return 244, 164, 96, 255 }
func BURLYWOODǁ2() (uint8, uint8, uint8, uint8)            { return 222, 184, 135, 255 }
func TANǁ2() (uint8, uint8, uint8, uint8)                  { return 210, 180, 140, 255 }
func ROSYBROWNǁ2() (uint8, uint8, uint8, uint8)            { return 188, 143, 143, 255 }
func MOCCASINǁ2() (uint8, uint8, uint8, uint8)             { return 255, 228, 181, 255 }
func NAVAJOWHITEǁ2() (uint8, uint8, uint8, uint8)          { return 255, 222, 173, 255 }
func PEACHPUFFǁ2() (uint8, uint8, uint8, uint8)            { return 255, 218, 185, 255 }
func MISTYROSEǁ2() (uint8, uint8, uint8, uint8)            { return 255, 228, 225, 255 }
func LAVENDERBLUSHǁ2() (uint8, uint8, uint8, uint8)        { return 255, 240, 245, 255 }
func LINENǁ2() (uint8, uint8, uint8, uint8)                { return 250, 240, 230, 255 }
func OLDLACEǁ2() (uint8, uint8, uint8, uint8)              { return 253, 245, 230, 255 }
func PAPAYAWHIPǁ2() (uint8, uint8, uint8, uint8)           { return 255, 239, 213, 255 }
func SEASHELLǁ2() (uint8, uint8, uint8, uint8)             { return 255, 245, 238, 255 }
func MINTCREAMǁ2() (uint8, uint8, uint8, uint8)            { return 245, 255, 250, 255 }
func SLATEGRAYǁ2() (uint8, uint8, uint8, uint8)            { return 112, 128, 144, 255 }
func LIGHTSLATEGRAYǁ2() (uint8, uint8, uint8, uint8)       { return 119, 136, 153, 255 }
func LIGHTSTEELBLUEǁ2() (uint8, uint8, uint8, uint8)       { return 176, 196, 222, 255 }
func LAVENDERǁ2() (uint8, uint8, uint8, uint8)             { return 230, 230, 250, 255 }
func FLORALWHITEǁ2() (uint8, uint8, uint8, uint8)          { return 255, 250, 240, 255 }
func ALICEBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 240, 248, 255, 255 }
func GHOSTWHITEǁ2() (uint8, uint8, uint8, uint8)           { return 248, 248, 255, 255 }
func HONEYDEWǁ2() (uint8, uint8, uint8, uint8)             { return 240, 255, 240, 255 }
func IVORYǁ2() (uint8, uint8, uint8, uint8)                { return 255, 255, 240, 255 }
func AZUREǁ2() (uint8, uint8, uint8, uint8)                { return 240, 255, 255, 255 }
func SNOWǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 250, 250, 255 }
func BLACKǁ2() (uint8, uint8, uint8, uint8)                { return 0, 0, 0, 255 }
func DIMGREYǁ2() (uint8, uint8, uint8, uint8)              { return 105, 105, 105, 255 }
func GREYǁ2() (uint8, uint8, uint8, uint8)                 { return 128, 128, 128, 255 }
func DARKGREYǁ2() (uint8, uint8, uint8, uint8)             { return 169, 169, 169, 255 }
func SILVERǁ2() (uint8, uint8, uint8, uint8)               { return 192, 192, 192, 255 }
func LIGHTGREYǁ2() (uint8, uint8, uint8, uint8)            { return 211, 211, 211, 255 }
func GAINSBOROǁ2() (uint8, uint8, uint8, uint8)            { return 220, 220, 220, 255 }
func WHITESMOKEǁ2() (uint8, uint8, uint8, uint8)           { return 245, 245, 245, 255 }
func WHITEǁ2() (uint8, uint8, uint8, uint8)                { return 255, 255, 255, 255 }

func MAROONǁ3() (rgba []uint8)         { return []uint8{128, 0, 0, 255} }
func DARKREDǁ3() (rgba []uint8)        { return []uint8{139, 0, 0, 255} }
func BROWNǁ3() (rgba []uint8)          { return []uint8{165, 42, 42, 255} }
func FIREBRICKǁ3() (rgba []uint8)      { return []uint8{178, 34, 34, 255} }
func CRIMSONǁ3() (rgba []uint8)        { return []uint8{220, 20, 60, 255} }
func REDǁ3() (rgba []uint8)            { return []uint8{255, 0, 0, 255} }
func TOMATOǁ3() (rgba []uint8)         { return []uint8{255, 99, 71, 255} }
func CORALǁ3() (rgba []uint8)          { return []uint8{255, 127, 80, 255} }
func INDIANREDǁ3() (rgba []uint8)      { return []uint8{205, 92, 92, 255} }
func LIGHTCORALǁ3() (rgba []uint8)     { return []uint8{240, 128, 128, 255} }
func DARKSALMONǁ3() (rgba []uint8)     { return []uint8{233, 150, 122, 255} }
func SALMONǁ3() (rgba []uint8)         { return []uint8{250, 128, 114, 255} }
func LIGHTSALMONǁ3() (rgba []uint8)    { return []uint8{255, 160, 122, 255} }
func ORANGEREDǁ3() (rgba []uint8)      { return []uint8{255, 69, 0, 255} }
func DARKORANGEǁ3() (rgba []uint8)     { return []uint8{255, 140, 0, 255} }
func ORANGEǁ3() (rgba []uint8)         { return []uint8{255, 165, 0, 255} }
func GOLDǁ3() (rgba []uint8)           { return []uint8{255, 215, 0, 255} }
func DARKGOLDENRODǁ3() (rgba []uint8)  { return []uint8{184, 134, 11, 255} }
func GOLDENRODǁ3() (rgba []uint8)      { return []uint8{218, 165, 32, 255} }
func PALEGOLDENRODǁ3() (rgba []uint8)  { return []uint8{238, 232, 170, 255} }
func DARKKHAKIǁ3() (rgba []uint8)      { return []uint8{189, 183, 107, 255} }
func KHAKIǁ3() (rgba []uint8)          { return []uint8{240, 230, 140, 255} }
func OLIVEǁ3() (rgba []uint8)          { return []uint8{128, 128, 0, 255} }
func YELLOWǁ3() (rgba []uint8)         { return []uint8{255, 255, 0, 255} }
func YELLOWGREENǁ3() (rgba []uint8)    { return []uint8{154, 205, 50, 255} }
func DARKOLIVEGREENǁ3() (rgba []uint8) { return []uint8{85, 107, 47, 255} }
func OLIVEDRABǁ3() (rgba []uint8)      { return []uint8{107, 142, 35, 255} }
func LAWNGREENǁ3() (rgba []uint8)      { return []uint8{124, 252, 0, 255} }
func CHARTREUSEǁ3() (rgba []uint8)     { return []uint8{127, 255, 0, 255} }
func GREENYELLOWǁ3() (rgba []uint8)    { return []uint8{173, 255, 47, 255} }
func DARKGREENǁ3() (rgba []uint8)      { return []uint8{0, 100, 0, 255} }
func GREENǁ3() (rgba []uint8)          { return []uint8{0, 128, 0, 255} }
func FORESTGREENǁ3() (rgba []uint8)    { return []uint8{34, 139, 34, 255} }
func LIMEǁ3() (rgba []uint8)           { return []uint8{0, 255, 0, 255} }
func LIMEGREENǁ3() (rgba []uint8)      { return []uint8{50, 205, 50, 255} }
func LIGHTGREENǁ3() (rgba []uint8)     { return []uint8{144, 238, 144, 255} }
func PALEGREENǁ3() (rgba []uint8)      { return []uint8{152, 251, 152, 255} }
func DARKSEAGREENǁ3() (rgba []uint8)   { return []uint8{143, 188, 143, 255} }
func MEDIUMSPRINGGREENǁ3() (rgba []uint8) {
	return []uint8{0, 250, 154, 255}
}
func SPRINGGREENǁ3() (rgba []uint8) { return []uint8{0, 255, 127, 255} }
func SEAGREENǁ3() (rgba []uint8)    { return []uint8{46, 139, 87, 255} }
func MEDIUMAQUAMARINEǁ3() (rgba []uint8) {
	return []uint8{102, 205, 170, 255}
}
func MEDIUMSEAGREENǁ3() (rgba []uint8)  { return []uint8{60, 179, 113, 255} }
func LIGHTSEAGREENǁ3() (rgba []uint8)   { return []uint8{32, 178, 170, 255} }
func DARKSLATEGRAYǁ3() (rgba []uint8)   { return []uint8{47, 79, 79, 255} }
func TEALǁ3() (rgba []uint8)            { return []uint8{0, 128, 128, 255} }
func DARKCYANǁ3() (rgba []uint8)        { return []uint8{0, 139, 139, 255} }
func AQUAǁ3() (rgba []uint8)            { return []uint8{0, 255, 255, 255} }
func CYANǁ3() (rgba []uint8)            { return []uint8{0, 255, 255, 255} }
func LIGHTCYANǁ3() (rgba []uint8)       { return []uint8{224, 255, 255, 255} }
func DARKTURQUOISEǁ3() (rgba []uint8)   { return []uint8{0, 206, 209, 255} }
func TURQUOISEǁ3() (rgba []uint8)       { return []uint8{64, 224, 208, 255} }
func MEDIUMTURQUOISEǁ3() (rgba []uint8) { return []uint8{72, 209, 204, 255} }
func PALETURQUOISEǁ3() (rgba []uint8)   { return []uint8{175, 238, 238, 255} }
func AQUAMARINEǁ3() (rgba []uint8)      { return []uint8{127, 255, 212, 255} }
func POWDERBLUEǁ3() (rgba []uint8)      { return []uint8{176, 224, 230, 255} }
func CADETBLUEǁ3() (rgba []uint8)       { return []uint8{95, 158, 160, 255} }
func STEELBLUEǁ3() (rgba []uint8)       { return []uint8{70, 130, 180, 255} }
func CORNFLOWERBLUEǁ3() (rgba []uint8)  { return []uint8{100, 149, 237, 255} }
func DEEPSKYBLUEǁ3() (rgba []uint8)     { return []uint8{0, 191, 255, 255} }
func DODGERBLUEǁ3() (rgba []uint8)      { return []uint8{30, 144, 255, 255} }
func LIGHTBLUEǁ3() (rgba []uint8)       { return []uint8{173, 216, 230, 255} }
func SKYBLUEǁ3() (rgba []uint8)         { return []uint8{135, 206, 235, 255} }
func LIGHTSKYBLUEǁ3() (rgba []uint8)    { return []uint8{135, 206, 250, 255} }
func MIDNIGHTBLUEǁ3() (rgba []uint8)    { return []uint8{25, 25, 112, 255} }
func NAVYǁ3() (rgba []uint8)            { return []uint8{0, 0, 128, 255} }
func DARKBLUEǁ3() (rgba []uint8)        { return []uint8{0, 0, 139, 255} }
func MEDIUMBLUEǁ3() (rgba []uint8)      { return []uint8{0, 0, 205, 255} }
func BLUEǁ3() (rgba []uint8)            { return []uint8{0, 0, 255, 255} }
func ROYALBLUEǁ3() (rgba []uint8)       { return []uint8{65, 105, 225, 255} }
func BLUEVIOLETǁ3() (rgba []uint8)      { return []uint8{138, 43, 226, 255} }
func INDIGOǁ3() (rgba []uint8)          { return []uint8{75, 0, 130, 255} }
func DARKSLATEBLUEǁ3() (rgba []uint8)   { return []uint8{72, 61, 139, 255} }
func SLATEBLUEǁ3() (rgba []uint8)       { return []uint8{106, 90, 205, 255} }
func MEDIUMSLATEBLUEǁ3() (rgba []uint8) {
	return []uint8{123, 104, 238, 255}
}
func MEDIUMPURPLEǁ3() (rgba []uint8)    { return []uint8{147, 112, 219, 255} }
func DARKMAGENTAǁ3() (rgba []uint8)     { return []uint8{139, 0, 139, 255} }
func DARKVIOLETǁ3() (rgba []uint8)      { return []uint8{148, 0, 211, 255} }
func DARKORCHIDǁ3() (rgba []uint8)      { return []uint8{153, 50, 204, 255} }
func MEDIUMORCHIDǁ3() (rgba []uint8)    { return []uint8{186, 85, 211, 255} }
func PURPLEǁ3() (rgba []uint8)          { return []uint8{128, 0, 128, 255} }
func THISTLEǁ3() (rgba []uint8)         { return []uint8{216, 191, 216, 255} }
func PLUMǁ3() (rgba []uint8)            { return []uint8{221, 160, 221, 255} }
func VIOLETǁ3() (rgba []uint8)          { return []uint8{238, 130, 238, 255} }
func MAGENTAǁ3() (rgba []uint8)         { return []uint8{255, 0, 255, 255} }
func ORCHIDǁ3() (rgba []uint8)          { return []uint8{218, 112, 214, 255} }
func MEDIUMVIOLETREDǁ3() (rgba []uint8) { return []uint8{199, 21, 133, 255} }
func PALEVIOLETREDǁ3() (rgba []uint8)   { return []uint8{219, 112, 147, 255} }
func DEEPPINKǁ3() (rgba []uint8)        { return []uint8{255, 20, 147, 255} }
func HOTPINKǁ3() (rgba []uint8)         { return []uint8{255, 105, 180, 255} }
func LIGHTPINKǁ3() (rgba []uint8)       { return []uint8{255, 182, 193, 255} }
func PINKǁ3() (rgba []uint8)            { return []uint8{255, 192, 203, 255} }
func ANTIQUEWHITEǁ3() (rgba []uint8)    { return []uint8{250, 235, 215, 255} }
func BEIGEǁ3() (rgba []uint8)           { return []uint8{245, 245, 220, 255} }
func BISQUEǁ3() (rgba []uint8)          { return []uint8{255, 228, 196, 255} }
func BLANCHEDALMONDǁ3() (rgba []uint8)  { return []uint8{255, 235, 205, 255} }
func WHEATǁ3() (rgba []uint8)           { return []uint8{245, 222, 179, 255} }
func CORNSILKǁ3() (rgba []uint8)        { return []uint8{255, 248, 220, 255} }
func LEMONCHIFFONǁ3() (rgba []uint8)    { return []uint8{255, 250, 205, 255} }
func LIGHTGOLDENRODYELLOWǁ3() (rgba []uint8) {
	return []uint8{250, 250, 210, 255}
}
func LIGHTYELLOWǁ3() (rgba []uint8)    { return []uint8{255, 255, 224, 255} }
func SADDLEBROWNǁ3() (rgba []uint8)    { return []uint8{139, 69, 19, 255} }
func SIENNAǁ3() (rgba []uint8)         { return []uint8{160, 82, 45, 255} }
func CHOCOLATEǁ3() (rgba []uint8)      { return []uint8{210, 105, 30, 255} }
func PERUǁ3() (rgba []uint8)           { return []uint8{205, 133, 63, 255} }
func SANDYBROWNǁ3() (rgba []uint8)     { return []uint8{244, 164, 96, 255} }
func BURLYWOODǁ3() (rgba []uint8)      { return []uint8{222, 184, 135, 255} }
func TANǁ3() (rgba []uint8)            { return []uint8{210, 180, 140, 255} }
func ROSYBROWNǁ3() (rgba []uint8)      { return []uint8{188, 143, 143, 255} }
func MOCCASINǁ3() (rgba []uint8)       { return []uint8{255, 228, 181, 255} }
func NAVAJOWHITEǁ3() (rgba []uint8)    { return []uint8{255, 222, 173, 255} }
func PEACHPUFFǁ3() (rgba []uint8)      { return []uint8{255, 218, 185, 255} }
func MISTYROSEǁ3() (rgba []uint8)      { return []uint8{255, 228, 225, 255} }
func LAVENDERBLUSHǁ3() (rgba []uint8)  { return []uint8{255, 240, 245, 255} }
func LINENǁ3() (rgba []uint8)          { return []uint8{250, 240, 230, 255} }
func OLDLACEǁ3() (rgba []uint8)        { return []uint8{253, 245, 230, 255} }
func PAPAYAWHIPǁ3() (rgba []uint8)     { return []uint8{255, 239, 213, 255} }
func SEASHELLǁ3() (rgba []uint8)       { return []uint8{255, 245, 238, 255} }
func MINTCREAMǁ3() (rgba []uint8)      { return []uint8{245, 255, 250, 255} }
func SLATEGRAYǁ3() (rgba []uint8)      { return []uint8{112, 128, 144, 255} }
func LIGHTSLATEGRAYǁ3() (rgba []uint8) { return []uint8{119, 136, 153, 255} }
func LIGHTSTEELBLUEǁ3() (rgba []uint8) { return []uint8{176, 196, 222, 255} }
func LAVENDERǁ3() (rgba []uint8)       { return []uint8{230, 230, 250, 255} }
func FLORALWHITEǁ3() (rgba []uint8)    { return []uint8{255, 250, 240, 255} }
func ALICEBLUEǁ3() (rgba []uint8)      { return []uint8{240, 248, 255, 255} }
func GHOSTWHITEǁ3() (rgba []uint8)     { return []uint8{248, 248, 255, 255} }
func HONEYDEWǁ3() (rgba []uint8)       { return []uint8{240, 255, 240, 255} }
func IVORYǁ3() (rgba []uint8)          { return []uint8{255, 255, 240, 255} }
func AZUREǁ3() (rgba []uint8)          { return []uint8{240, 255, 255, 255} }
func SNOWǁ3() (rgba []uint8)           { return []uint8{255, 250, 250, 255} }
func BLACKǁ3() (rgba []uint8)          { return []uint8{0, 0, 0, 255} }
func DIMGREYǁ3() (rgba []uint8)        { return []uint8{105, 105, 105, 255} }
func GREYǁ3() (rgba []uint8)           { return []uint8{128, 128, 128, 255} }
func DARKGREYǁ3() (rgba []uint8)       { return []uint8{169, 169, 169, 255} }
func SILVERǁ3() (rgba []uint8)         { return []uint8{192, 192, 192, 255} }
func LIGHTGREYǁ3() (rgba []uint8)      { return []uint8{211, 211, 211, 255} }
func GAINSBOROǁ3() (rgba []uint8)      { return []uint8{220, 220, 220, 255} }
func WHITESMOKEǁ3() (rgba []uint8)     { return []uint8{245, 245, 245, 255} }
func WHITEǁ3() (rgba []uint8)          { return []uint8{255, 255, 255, 255} }
