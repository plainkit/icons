package lucide

import x "github.com/bloxui/blox"

// Codepen creates a Codepen Lucide icon.
func Codepen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-codepen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polygon(x.Points("12 2 22 8.5 22 15.5 12 22 2 15.5 2 8.5 12 2"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("22"), x.Y2("15.5"))),
		x.Child(x.Polyline(x.Points("22 8.5 12 15.5 2 8.5"))),
		x.Child(x.Polyline(x.Points("2 15.5 12 8.5 22 15.5"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("2"), x.Y2("8.5"))),
	)
	return x.Svg(svgArgs...)
}
