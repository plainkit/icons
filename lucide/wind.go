package lucide

import (
	html "github.com/plainkit/html"
)

// Wind creates a Wind Lucide icon.
func Wind(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wind", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.8 19.6A2 2 0 1 0 14 16H2"))),
		html.Child(html.SvgPath(html.AD("M17.5 8a2.5 2.5 0 1 1 2 4H2"))),
		html.Child(html.SvgPath(html.AD("M9.8 4.4A2 2 0 1 1 11 8H2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
