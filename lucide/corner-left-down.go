package lucide

import (
	html "github.com/plainkit/html"
)

// CornerLeftDown creates a Corner Left Down Lucide icon.
func CornerLeftDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-left-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 15-5 5-5-5")),
		html.SvgPath(html.AD("M20 4h-7a4 4 0 0 0-4 4v12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
