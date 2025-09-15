package lucide

import x "github.com/plainkit/html"

// SquareArrowDown creates a Square Arrow Down Lucide icon.
func SquareArrowDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 8v8"))),
		x.Child(x.Path(x.D("m8 12 4 4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
