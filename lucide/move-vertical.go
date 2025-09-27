package lucide

import (
	html "github.com/plainkit/html"
)

// MoveVertical creates a Move Vertical Lucide icon.
func MoveVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v20"))),
		html.Child(html.SvgPath(html.AD("m8 18 4 4 4-4"))),
		html.Child(html.SvgPath(html.AD("m8 6 4-4 4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
