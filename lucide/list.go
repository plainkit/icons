package lucide

import (
	html "github.com/plainkit/html"
)

// List creates a List Lucide icon.
func List(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 5h.01")),
		html.SvgPath(html.AD("M3 12h.01")),
		html.SvgPath(html.AD("M3 19h.01")),
		html.SvgPath(html.AD("M8 5h13")),
		html.SvgPath(html.AD("M8 12h13")),
		html.SvgPath(html.AD("M8 19h13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
