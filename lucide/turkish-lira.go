package lucide

import (
	html "github.com/plainkit/html"
)

// TurkishLira creates a Turkish Lira Lucide icon.
func TurkishLira(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-turkish-lira", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 4 5 9")),
		html.SvgPath(html.AD("m15 8.5-10 5")),
		html.SvgPath(html.AD("M18 12a9 9 0 0 1-9 9V3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
