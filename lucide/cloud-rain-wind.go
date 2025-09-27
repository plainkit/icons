package lucide

import (
	html "github.com/plainkit/html"
)

// CloudRainWind creates a Cloud Rain Wind Lucide icon.
func CloudRainWind(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-rain-wind", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		html.Child(html.SvgPath(html.AD("m9.2 22 3-7"))),
		html.Child(html.SvgPath(html.AD("m9 13-3 7"))),
		html.Child(html.SvgPath(html.AD("m17 13-3 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
