package lucide

import x "github.com/bloxui/blox"

// Sword creates a Sword Lucide icon.
func Sword(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-sword", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polyline(x.Points("14.5 17.5 3 6 3 3 6 3 17.5 14.5"))),
		x.Child(x.Line(x.X1("13"), x.X2("19"), x.Y1("19"), x.Y2("13"))),
		x.Child(x.Line(x.X1("16"), x.X2("20"), x.Y1("16"), x.Y2("20"))),
		x.Child(x.Line(x.X1("19"), x.X2("21"), x.Y1("21"), x.Y2("19"))),
	)
	return x.Svg(svgArgs...)
}
