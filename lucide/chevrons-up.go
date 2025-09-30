package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsUp creates a Chevrons Up Lucide icon.
func ChevronsUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m17 11-5-5-5 5")),
		html.SvgPath(html.AD("m17 18-5-5-5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
