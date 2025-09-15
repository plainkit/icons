package lucide

import x "github.com/plainkit/blox"

// ZoomOut creates a Zoom Out Lucide icon.
func ZoomOut(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-zoom-out", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
		x.Child(x.Line(x.X1("21"), x.X2("16.65"), x.Y1("21"), x.Y2("16.65"))),
		x.Child(x.Line(x.X1("8"), x.X2("14"), x.Y1("11"), x.Y2("11"))),
	)
	return x.Svg(svgArgs...)
}
