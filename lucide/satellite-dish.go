package lucide

import (
	html "github.com/plainkit/html"
)

// SatelliteDish creates a Satellite Dish Lucide icon.
func SatelliteDish(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-satellite-dish", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 10a7.31 7.31 0 0 0 10 10Z")),
		html.SvgPath(html.AD("m9 15 3-3")),
		html.SvgPath(html.AD("M17 13a6 6 0 0 0-6-6")),
		html.SvgPath(html.AD("M21 13A10 10 0 0 0 11 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
