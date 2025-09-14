package lucide

import x "github.com/bloxui/blox"

// SquareChevronDown creates a Square Chevron Down Lucide icon.
func SquareChevronDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-chevron-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m16 10-4 4-4-4"))),
	)
	return x.Svg(svgArgs...)
}
