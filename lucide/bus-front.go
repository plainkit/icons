package lucide

import x "github.com/bloxui/blox"

// BusFront creates a Bus Front Lucide icon.
func BusFront(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bus-front", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 6 2 7"))),
		x.Child(x.Path(x.D("M10 6h4"))),
		x.Child(x.Path(x.D("m22 7-2-1"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("16"), x.X("4"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 11h16"))),
		x.Child(x.Path(x.D("M8 15h.01"))),
		x.Child(x.Path(x.D("M16 15h.01"))),
		x.Child(x.Path(x.D("M6 19v2"))),
		x.Child(x.Path(x.D("M18 21v-2"))),
	)
	return x.Svg(svgArgs...)
}
