package lucide

import x "github.com/plainkit/blox"

// SquareArrowDownRight creates a Square Arrow Down Right Lucide icon.
func SquareArrowDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m8 8 8 8"))),
		x.Child(x.Path(x.D("M16 8v8H8"))),
	)
	return x.Svg(svgArgs...)
}
