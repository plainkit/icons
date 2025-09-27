package lucide

import (
	html "github.com/plainkit/html"
)

// ListChevronsDownUp creates a List Chevrons Down Up Lucide icon.
func ListChevronsDownUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-chevrons-down-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 5h8"))),
		html.Child(html.SvgPath(html.AD("M3 12h8"))),
		html.Child(html.SvgPath(html.AD("M3 19h8"))),
		html.Child(html.SvgPath(html.AD("m15 5 3 3 3-3"))),
		html.Child(html.SvgPath(html.AD("m15 19 3-3 3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
