package lucide

import (
	html "github.com/plainkit/html"
)

// MoveDown creates a Move Down Lucide icon.
func MoveDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 18L12 22L16 18"))),
		html.Child(html.SvgPath(html.AD("M12 2V22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
