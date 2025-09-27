package lucide

import (
	html "github.com/plainkit/html"
)

// MoveRight creates a Move Right Lucide icon.
func MoveRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 8L22 12L18 16"))),
		html.Child(html.SvgPath(html.AD("M2 12H22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
