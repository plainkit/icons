package lucide

import x "github.com/plainkit/html"

// SquareArrowUpLeft creates a Square Arrow Up Left Lucide icon.
func SquareArrowUpLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 16V8h8"))),
		x.Child(x.Path(x.D("M16 16 8 8"))),
	)
	return x.Svg(svgArgs...)
}
