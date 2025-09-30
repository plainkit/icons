package lucide

import (
	html "github.com/plainkit/html"
)

// ListX creates a List X Lucide icon.
func ListX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 5H3")),
		html.SvgPath(html.AD("M11 12H3")),
		html.SvgPath(html.AD("M16 19H3")),
		html.SvgPath(html.AD("m15.5 9.5 5 5")),
		html.SvgPath(html.AD("m20.5 9.5-5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
