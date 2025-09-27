package lucide

import (
	html "github.com/plainkit/html"
)

// MoveHorizontal creates a Move Horizontal Lucide icon.
func MoveHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 8 4 4-4 4"))),
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
		html.Child(html.SvgPath(html.AD("m6 8-4 4 4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
