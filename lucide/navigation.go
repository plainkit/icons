package lucide

import (
	html "github.com/plainkit/html"
)

// Navigation creates a Navigation Lucide icon.
func Navigation(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-navigation", args)
	children := []html.SvgArg{
		html.Child(html.SvgPolygon(html.APoints("3 11 22 2 13 21 11 13 3 11"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
