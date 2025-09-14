package lucide

import x "github.com/bloxui/blox"

// Navigation creates a Navigation Lucide icon.
func Navigation(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-navigation", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polygon(x.Points("3 11 22 2 13 21 11 13 3 11"))),
	)
	return x.Svg(svgArgs...)
}
