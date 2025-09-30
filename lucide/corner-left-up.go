package lucide

import (
	html "github.com/plainkit/html"
)

// CornerLeftUp creates a Corner Left Up Lucide icon.
func CornerLeftUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-left-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 9 9 4 4 9")),
		html.SvgPath(html.AD("M20 20h-7a4 4 0 0 1-4-4V4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
