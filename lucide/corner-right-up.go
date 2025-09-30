package lucide

import (
	html "github.com/plainkit/html"
)

// CornerRightUp creates a Corner Right Up Lucide icon.
func CornerRightUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-right-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10 9 5-5 5 5")),
		html.SvgPath(html.AD("M4 20h7a4 4 0 0 0 4-4V4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
