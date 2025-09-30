package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsDown creates a Chevrons Down Lucide icon.
func ChevronsDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 6 5 5 5-5")),
		html.SvgPath(html.AD("m7 13 5 5 5-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
