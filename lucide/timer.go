package lucide

import x "github.com/bloxui/blox"

// Timer creates a Timer Lucide icon.
func Timer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-timer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("10"), x.X2("14"), x.Y1("2"), x.Y2("2"))),
		x.Child(x.Line(x.X1("12"), x.X2("15"), x.Y1("14"), x.Y2("11"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("14"), x.R("8"))),
	)
	return x.Svg(svgArgs...)
}
