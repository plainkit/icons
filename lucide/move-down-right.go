package lucide

import (
	html "github.com/plainkit/html"
)

// MoveDownRight creates a Move Down Right Lucide icon.
func MoveDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-down-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19 13V19H13"))),
		html.Child(html.SvgPath(html.AD("M5 5L19 19"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
