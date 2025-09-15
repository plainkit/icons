package lucide

import x "github.com/plainkit/blox"

// Wind creates a Wind Lucide icon.
func Wind(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wind", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.8 19.6A2 2 0 1 0 14 16H2"))),
		x.Child(x.Path(x.D("M17.5 8a2.5 2.5 0 1 1 2 4H2"))),
		x.Child(x.Path(x.D("M9.8 4.4A2 2 0 1 1 11 8H2"))),
	)
	return x.Svg(svgArgs...)
}
