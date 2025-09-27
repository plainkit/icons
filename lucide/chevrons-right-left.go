package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsRightLeft creates a Chevrons Right Left Lucide icon.
func ChevronsRightLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-right-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m20 17-5-5 5-5"))),
		html.Child(html.SvgPath(html.AD("m4 17 5-5-5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
