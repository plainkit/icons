package lucide

import (
	html "github.com/plainkit/html"
)

// Pencil creates a Pencil Lucide icon.
func Pencil(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pencil", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z")),
		html.SvgPath(html.AD("m15 5 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
