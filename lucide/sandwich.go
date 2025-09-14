package lucide

import x "github.com/bloxui/blox"

// Sandwich creates a Sandwich Lucide icon.
func Sandwich(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sandwich", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2.37 11.223 8.372-6.777a2 2 0 0 1 2.516 0l8.371 6.777"))),
		x.Child(x.Path(x.D("M21 15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-5.25"))),
		x.Child(x.Path(x.D("M3 15a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h9"))),
		x.Child(x.Path(x.D("m6.67 15 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("4"), x.X("2"), x.Y("11"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
