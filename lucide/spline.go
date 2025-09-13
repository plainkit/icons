package lucide

import x "github.com/bloxui/blox"

// Spline creates a Spline Lucide icon.
func Spline(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-spline", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("19"), x.Cy("5"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("2"))),
		x.Child(x.Path(x.D("M5 17A12 12 0 0 1 17 5"))),
	)
	return x.Svg(svgArgs...)
}
