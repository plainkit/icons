package lucide

import x "github.com/plainkit/html"

// SquareCheck creates a Square Check Lucide icon.
func SquareCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m9 12 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
