package lucide

import x "github.com/bloxui/blox"

// Router creates a Router Lucide icon.
func Router(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-router", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("8"), x.X("2"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6.01 18H6"))),
		x.Child(x.Path(x.D("M10.01 18H10"))),
		x.Child(x.Path(x.D("M15 10v4"))),
		x.Child(x.Path(x.D("M17.84 7.17a4 4 0 0 0-5.66 0"))),
		x.Child(x.Path(x.D("M20.66 4.34a8 8 0 0 0-11.31 0"))),
	)
	return x.Svg(svgArgs...)
}
