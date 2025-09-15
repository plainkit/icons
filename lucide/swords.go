package lucide

import x "github.com/plainkit/html"

// Swords creates a Swords Lucide icon.
func Swords(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-swords", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polyline(x.Points("14.5 17.5 3 6 3 3 6 3 17.5 14.5"))),
		x.Child(x.Line(x.X1("13"), x.X2("19"), x.Y1("19"), x.Y2("13"))),
		x.Child(x.Line(x.X1("16"), x.X2("20"), x.Y1("16"), x.Y2("20"))),
		x.Child(x.Line(x.X1("19"), x.X2("21"), x.Y1("21"), x.Y2("19"))),
		x.Child(x.Polyline(x.Points("14.5 6.5 18 3 21 3 21 6 17.5 9.5"))),
		x.Child(x.Line(x.X1("5"), x.X2("9"), x.Y1("14"), x.Y2("18"))),
		x.Child(x.Line(x.X1("7"), x.X2("4"), x.Y1("17"), x.Y2("20"))),
		x.Child(x.Line(x.X1("3"), x.X2("5"), x.Y1("19"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
