package lucide

import x "github.com/plainkit/blox"

// Currency creates a Currency Lucide icon.
func Currency(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-currency", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("8"))),
		x.Child(x.Line(x.X1("3"), x.X2("6"), x.Y1("3"), x.Y2("6"))),
		x.Child(x.Line(x.X1("21"), x.X2("18"), x.Y1("3"), x.Y2("6"))),
		x.Child(x.Line(x.X1("3"), x.X2("6"), x.Y1("21"), x.Y2("18"))),
		x.Child(x.Line(x.X1("21"), x.X2("18"), x.Y1("21"), x.Y2("18"))),
	)
	return x.Svg(svgArgs...)
}
