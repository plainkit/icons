package lucide

import x "github.com/bloxui/blox"

// SquareChevronRight creates a Square Chevron Right Lucide icon.
func SquareChevronRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-chevron-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m10 8 4 4-4 4"))),
	)
	return x.Svg(svgArgs...)
}
