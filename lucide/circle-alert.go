package lucide

import x "github.com/plainkit/html"

// CircleAlert creates a Circle Alert Lucide icon.
func CircleAlert(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("8"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12.01"), x.Y1("16"), x.Y2("16"))),
	)
	return x.Svg(svgArgs...)
}
