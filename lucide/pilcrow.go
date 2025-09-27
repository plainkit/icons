package lucide

import (
	html "github.com/plainkit/html"
)

// Pilcrow creates a Pilcrow Lucide icon.
func Pilcrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pilcrow", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 4v16"))),
		html.Child(html.SvgPath(html.AD("M17 4v16"))),
		html.Child(html.SvgPath(html.AD("M19 4H9.5a4.5 4.5 0 0 0 0 9H13"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
