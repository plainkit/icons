package lucide

import (
	html "github.com/plainkit/html"
)

// CornerDownRight creates a Corner Down Right Lucide icon.
func CornerDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-down-right", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 10 5 5-5 5")),
		html.SvgPath(html.AD("M4 4v7a4 4 0 0 0 4 4h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
