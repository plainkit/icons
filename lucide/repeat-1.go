package lucide

import (
	html "github.com/plainkit/html"
)

// Repeat1 creates a Repeat 1 Lucide icon.
func Repeat1(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-repeat-1", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m17 2 4 4-4 4")),
		html.SvgPath(html.AD("M3 11v-1a4 4 0 0 1 4-4h14")),
		html.SvgPath(html.AD("m7 22-4-4 4-4")),
		html.SvgPath(html.AD("M21 13v1a4 4 0 0 1-4 4H3")),
		html.SvgPath(html.AD("M11 10h1v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
