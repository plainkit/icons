package lucide

import x "github.com/plainkit/html"

// Navigation2 creates a Navigation 2 Lucide icon.
func Navigation2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-navigation-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polygon(x.Points("12 2 19 21 12 17 5 21 12 2"))),
	)
	return x.Svg(svgArgs...)
}
