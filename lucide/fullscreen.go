package lucide

import x "github.com/bloxui/blox"

// Fullscreen creates a Fullscreen Lucide icon.
func Fullscreen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fullscreen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("8"), x.X("7"), x.Y("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
