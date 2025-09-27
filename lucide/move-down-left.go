package lucide

import (
	html "github.com/plainkit/html"
)

// MoveDownLeft creates a Move Down Left Lucide icon.
func MoveDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-down-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 19H5V13"))),
		html.Child(html.SvgPath(html.AD("M19 5L5 19"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
