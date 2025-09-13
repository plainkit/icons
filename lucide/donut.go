package lucide

import x "github.com/bloxui/blox"

// Donut creates a Donut Lucide icon.
func Donut(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-donut", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4 10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
