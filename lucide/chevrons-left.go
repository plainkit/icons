package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsLeft creates a Chevrons Left Lucide icon.
func ChevronsLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m11 17-5-5 5-5")),
		html.SvgPath(html.AD("m18 17-5-5 5-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
