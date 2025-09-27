package lucide

import (
	html "github.com/plainkit/html"
)

// Repeat creates a Repeat Lucide icon.
func Repeat(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-repeat", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m17 2 4 4-4 4"))),
		html.Child(html.SvgPath(html.AD("M3 11v-1a4 4 0 0 1 4-4h14"))),
		html.Child(html.SvgPath(html.AD("m7 22-4-4 4-4"))),
		html.Child(html.SvgPath(html.AD("M21 13v1a4 4 0 0 1-4 4H3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
