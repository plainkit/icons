package lucide

import x "github.com/plainkit/html"

// Tablet creates a Tablet Lucide icon.
func Tablet(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tablet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Line(x.X1("12"), x.X2("12.01"), x.Y1("18"), x.Y2("18"))),
	)
	return x.Svg(svgArgs...)
}
