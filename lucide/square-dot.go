package lucide

import x "github.com/plainkit/html"

// SquareDot creates a Square Dot Lucide icon.
func SquareDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
