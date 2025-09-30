package lucide

import (
	html "github.com/plainkit/html"
)

// Navigation2 creates a Navigation 2 Lucide icon.
func Navigation2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-navigation-2", args)
	children := []html.SvgArg{
		html.SvgPolygon(html.APoints("12 2 19 21 12 17 5 21 12 2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
