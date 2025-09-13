package lucide

import x "github.com/bloxui/blox"

// Crosshair creates a Crosshair Lucide icon.
func Crosshair(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-crosshair", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Line(x.X1("22"), x.X2("18"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("6"), x.X2("2"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("6"), x.Y2("2"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("22"), x.Y2("18"))),
	)
	return x.Svg(svgArgs...)
}