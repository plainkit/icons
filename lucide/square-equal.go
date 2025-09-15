package lucide

import x "github.com/plainkit/html"

// SquareEqual creates a Square Equal Lucide icon.
func SquareEqual(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-equal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 10h10"))),
		x.Child(x.Path(x.D("M7 14h10"))),
	)
	return x.Svg(svgArgs...)
}
