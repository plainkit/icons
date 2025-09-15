package lucide

import x "github.com/plainkit/html"

// SquareArrowUp creates a Square Arrow Up Lucide icon.
func SquareArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m16 12-4-4-4 4"))),
		x.Child(x.Path(x.D("M12 16V8"))),
	)
	return x.Svg(svgArgs...)
}
