package lucide

import x "github.com/bloxui/blox"

// Meh creates a Meh Lucide icon.
func Meh(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-meh", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("15"), x.Y2("15"))),
		x.Child(x.Line(x.X1("9"), x.X2("9.01"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("9"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
