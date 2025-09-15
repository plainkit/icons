package lucide

import x "github.com/plainkit/blox"

// SlidersHorizontal creates a Sliders Horizontal Lucide icon.
func SlidersHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sliders-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 5H3"))),
		x.Child(x.Path(x.D("M12 19H3"))),
		x.Child(x.Path(x.D("M14 3v4"))),
		x.Child(x.Path(x.D("M16 17v4"))),
		x.Child(x.Path(x.D("M21 12h-9"))),
		x.Child(x.Path(x.D("M21 19h-5"))),
		x.Child(x.Path(x.D("M21 5h-7"))),
		x.Child(x.Path(x.D("M8 10v4"))),
		x.Child(x.Path(x.D("M8 12H3"))),
	)
	return x.Svg(svgArgs...)
}
