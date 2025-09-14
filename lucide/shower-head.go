package lucide

import x "github.com/bloxui/blox"

// ShowerHead creates a Shower Head Lucide icon.
func ShowerHead(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shower-head", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m4 4 2.5 2.5"))),
		x.Child(x.Path(x.D("M13.5 6.5a4.95 4.95 0 0 0-7 7"))),
		x.Child(x.Path(x.D("M15 5 5 15"))),
		x.Child(x.Path(x.D("M14 17v.01"))),
		x.Child(x.Path(x.D("M10 16v.01"))),
		x.Child(x.Path(x.D("M13 13v.01"))),
		x.Child(x.Path(x.D("M16 10v.01"))),
		x.Child(x.Path(x.D("M11 20v.01"))),
		x.Child(x.Path(x.D("M17 14v.01"))),
		x.Child(x.Path(x.D("M20 11v.01"))),
	)
	return x.Svg(svgArgs...)
}
