package lucide

import x "github.com/bloxui/blox"

// SquareArrowRight creates a Square Arrow Right Lucide icon.
func SquareArrowRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("m12 16 4-4-4-4"))),
	)
	return x.Svg(svgArgs...)
}
