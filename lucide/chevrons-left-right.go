package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsLeftRight creates a Chevrons Left Right Lucide icon.
func ChevronsLeftRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-left-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m9 7-5 5 5 5"))),
		html.Child(html.SvgPath(html.AD("m15 7 5 5-5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
