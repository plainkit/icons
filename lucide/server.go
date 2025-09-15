package lucide

import x "github.com/plainkit/blox"

// Server creates a Server Lucide icon.
func Server(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-server", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("8"), x.X("2"), x.Y("2"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("8"), x.X("2"), x.Y("14"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Line(x.X1("6"), x.X2("6.01"), x.Y1("6"), x.Y2("6"))),
		x.Child(x.Line(x.X1("6"), x.X2("6.01"), x.Y1("18"), x.Y2("18"))),
	)
	return x.Svg(svgArgs...)
}
