package lucide

import x "github.com/plainkit/blox"

// SquareArrowLeft creates a Square Arrow Left Lucide icon.
func SquareArrowLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m12 8-4 4 4 4"))),
		x.Child(x.Path(x.D("M16 12H8"))),
	)
	return x.Svg(svgArgs...)
}
