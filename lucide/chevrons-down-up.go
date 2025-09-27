package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsDownUp creates a Chevrons Down Up Lucide icon.
func ChevronsDownUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-down-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m7 20 5-5 5 5"))),
		html.Child(html.SvgPath(html.AD("m7 4 5 5 5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
