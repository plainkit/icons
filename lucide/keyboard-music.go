package lucide

import x "github.com/bloxui/blox"

// KeyboardMusic creates a Keyboard Music Lucide icon.
func KeyboardMusic(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-keyboard-music", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 8h4"))),
		x.Child(x.Path(x.D("M14 8h.01"))),
		x.Child(x.Path(x.D("M18 8h.01"))),
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("M6 12v4"))),
		x.Child(x.Path(x.D("M10 12v4"))),
		x.Child(x.Path(x.D("M14 12v4"))),
		x.Child(x.Path(x.D("M18 12v4"))),
	)
	return x.Svg(svgArgs...)
}
