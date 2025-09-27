package lucide

import (
	html "github.com/plainkit/html"
)

// ListChevronsUpDown creates a List Chevrons Up Down Lucide icon.
func ListChevronsUpDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-chevrons-up-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 5h8"))),
		html.Child(html.SvgPath(html.AD("M3 12h8"))),
		html.Child(html.SvgPath(html.AD("M3 19h8"))),
		html.Child(html.SvgPath(html.AD("m15 8 3-3 3 3"))),
		html.Child(html.SvgPath(html.AD("m15 16 3 3 3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
