package lucide

import x "github.com/plainkit/html"

// Percent creates a Percent Lucide icon.
func Percent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("19"), x.X2("5"), x.Y1("5"), x.Y2("19"))),
		x.Child(x.Circle(x.Cx("6.5"), x.Cy("6.5"), x.R("2.5"))),
		x.Child(x.Circle(x.Cx("17.5"), x.Cy("17.5"), x.R("2.5"))),
	)
	return x.Svg(svgArgs...)
}
