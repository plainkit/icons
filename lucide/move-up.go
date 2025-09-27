package lucide

import (
	html "github.com/plainkit/html"
)

// MoveUp creates a Move Up Lucide icon.
func MoveUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 6L12 2L16 6"))),
		html.Child(html.SvgPath(html.AD("M12 2V22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
