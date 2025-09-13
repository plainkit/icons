package lucide

import x "github.com/bloxui/blox"

// Earth creates a Earth Lucide icon.
func Earth(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-earth", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.54 15H17a2 2 0 0 0-2 2v4.54"))),
		x.Child(x.Path(x.D("M7 3.34V5a3 3 0 0 0 3 3a2 2 0 0 1 2 2c0 1.1.9 2 2 2a2 2 0 0 0 2-2c0-1.1.9-2 2-2h3.17"))),
		x.Child(x.Path(x.D("M11 21.95V18a2 2 0 0 0-2-2a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}