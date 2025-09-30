package lucide

import (
	html "github.com/plainkit/html"
)

// Repeat2 creates a Repeat 2 Lucide icon.
func Repeat2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-repeat-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 9 3-3 3 3")),
		html.SvgPath(html.AD("M13 18H7a2 2 0 0 1-2-2V6")),
		html.SvgPath(html.AD("m22 15-3 3-3-3")),
		html.SvgPath(html.AD("M11 6h6a2 2 0 0 1 2 2v10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
