package lucide

import x "github.com/bloxui/blox"

// SlidersVertical creates a Sliders Vertical Lucide icon.
func SlidersVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sliders-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 8h4"))),
		x.Child(x.Path(x.D("M12 21v-9"))),
		x.Child(x.Path(x.D("M12 8V3"))),
		x.Child(x.Path(x.D("M17 16h4"))),
		x.Child(x.Path(x.D("M19 12V3"))),
		x.Child(x.Path(x.D("M19 21v-5"))),
		x.Child(x.Path(x.D("M3 14h4"))),
		x.Child(x.Path(x.D("M5 10V3"))),
		x.Child(x.Path(x.D("M5 21v-7"))),
	)
	return x.Svg(svgArgs...)
}
