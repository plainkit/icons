package lucide

import x "github.com/plainkit/blox"

// SquareChevronLeft creates a Square Chevron Left Lucide icon.
func SquareChevronLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-chevron-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m14 16-4-4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
