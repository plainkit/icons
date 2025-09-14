package lucide

import x "github.com/bloxui/blox"

// Divide creates a Divide Lucide icon.
func Divide(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-divide", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("6"), x.R("1"))),
		x.Child(x.Line(x.X1("5"), x.X2("19"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("18"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
