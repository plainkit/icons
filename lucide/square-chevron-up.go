package lucide

import x "github.com/plainkit/blox"

// SquareChevronUp creates a Square Chevron Up Lucide icon.
func SquareChevronUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-chevron-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m8 14 4-4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
