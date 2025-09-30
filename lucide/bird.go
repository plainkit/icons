package lucide

import (
	html "github.com/plainkit/html"
)

// Bird creates a Bird Lucide icon.
func Bird(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bird", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 7h.01")),
		html.SvgPath(html.AD("M3.4 18H12a8 8 0 0 0 8-8V7a4 4 0 0 0-7.28-2.3L2 20")),
		html.SvgPath(html.AD("m20 7 2 .5-2 .5")),
		html.SvgPath(html.AD("M10 18v3")),
		html.SvgPath(html.AD("M14 17.75V21")),
		html.SvgPath(html.AD("M7 18a6 6 0 0 0 3.84-10.61")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
