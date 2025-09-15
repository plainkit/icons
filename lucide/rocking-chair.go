package lucide

import x "github.com/plainkit/blox"

// RockingChair creates a Rocking Chair Lucide icon.
func RockingChair(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rocking-chair", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polyline(x.Points("3.5 2 6.5 12.5 18 12.5"))),
		x.Child(x.Line(x.X1("9.5"), x.X2("5.5"), x.Y1("12.5"), x.Y2("20"))),
		x.Child(x.Line(x.X1("15"), x.X2("18.5"), x.Y1("12.5"), x.Y2("20"))),
		x.Child(x.Path(x.D("M2.75 18a13 13 0 0 0 18.5 0"))),
	)
	return x.Svg(svgArgs...)
}
