package lucide

import (
	html "github.com/plainkit/html"
)

// MoveUpRight creates a Move Up Right Lucide icon.
func MoveUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-up-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 5H19V11"))),
		html.Child(html.SvgPath(html.AD("M19 5L5 19"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
