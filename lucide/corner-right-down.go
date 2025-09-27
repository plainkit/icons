package lucide

import (
	html "github.com/plainkit/html"
)

// CornerRightDown creates a Corner Right Down Lucide icon.
func CornerRightDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-right-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10 15 5 5 5-5"))),
		html.Child(html.SvgPath(html.AD("M4 4h7a4 4 0 0 1 4 4v12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
