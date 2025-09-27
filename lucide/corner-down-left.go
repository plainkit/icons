package lucide

import (
	html "github.com/plainkit/html"
)

// CornerDownLeft creates a Corner Down Left Lucide icon.
func CornerDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-down-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 4v7a4 4 0 0 1-4 4H4"))),
		html.Child(html.SvgPath(html.AD("m9 10-5 5 5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
