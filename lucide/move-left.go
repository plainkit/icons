package lucide

import (
	html "github.com/plainkit/html"
)

// MoveLeft creates a Move Left Lucide icon.
func MoveLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 8L2 12L6 16")),
		html.SvgPath(html.AD("M2 12H22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
