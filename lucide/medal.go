package lucide

import x "github.com/bloxui/blox"

// Medal creates a Medal Lucide icon.
func Medal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-medal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7.21 15 2.66 7.14a2 2 0 0 1 .13-2.2L4.4 2.8A2 2 0 0 1 6 2h12a2 2 0 0 1 1.6.8l1.6 2.14a2 2 0 0 1 .14 2.2L16.79 15"))),
		x.Child(x.Path(x.D("M11 12 5.12 2.2"))),
		x.Child(x.Path(x.D("m13 12 5.88-9.8"))),
		x.Child(x.Path(x.D("M8 7h8"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("17"), x.R("5"))),
		x.Child(x.Path(x.D("M12 18v-2h-.5"))),
	)
	return x.Svg(svgArgs...)
}
