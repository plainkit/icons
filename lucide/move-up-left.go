package lucide

import (
	html "github.com/plainkit/html"
)

// MoveUpLeft creates a Move Up Left Lucide icon.
func MoveUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-up-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 11V5H11"))),
		html.Child(html.SvgPath(html.AD("M5 5L19 19"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
