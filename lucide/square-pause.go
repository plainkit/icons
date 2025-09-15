package lucide

import x "github.com/plainkit/html"

// SquarePause creates a Square Pause Lucide icon.
func SquarePause(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-pause", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Line(x.X1("10"), x.X2("10"), x.Y1("15"), x.Y2("9"))),
		x.Child(x.Line(x.X1("14"), x.X2("14"), x.Y1("15"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
