package lucide

import x "github.com/bloxui/blox"

// Lasso creates a Lasso Lucide icon.
func Lasso(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-lasso", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.704 14.467A10 8 0 0 1 2 10a10 8 0 0 1 20 0 10 8 0 0 1-10 8 10 8 0 0 1-5.181-1.158"))),
		x.Child(x.Path(x.D("M7 22a5 5 0 0 1-2-3.994"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("16"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
