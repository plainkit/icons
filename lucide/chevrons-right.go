package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsRight creates a Chevrons Right Lucide icon.
func ChevronsRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m6 17 5-5-5-5"))),
		html.Child(html.SvgPath(html.AD("m13 17 5-5-5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
