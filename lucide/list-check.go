package lucide

import (
	html "github.com/plainkit/html"
)

// ListCheck creates a List Check Lucide icon.
func ListCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 5H3")),
		html.SvgPath(html.AD("M16 12H3")),
		html.SvgPath(html.AD("M11 19H3")),
		html.SvgPath(html.AD("m15 18 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
