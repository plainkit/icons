package lucide

import x "github.com/bloxui/blox"

// Route creates a Route Lucide icon.
func Route(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-route", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("19"), x.R("3"))),
		x.Child(x.Path(x.D("M9 19h8.5a3.5 3.5 0 0 0 0-7h-11a3.5 3.5 0 0 1 0-7H15"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("5"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
