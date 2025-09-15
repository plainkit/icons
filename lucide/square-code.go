package lucide

import x "github.com/plainkit/html"

// SquareCode creates a Square Code Lucide icon.
func SquareCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 9-3 3 3 3"))),
		x.Child(x.Path(x.D("m14 15 3-3-3-3"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
