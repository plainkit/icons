package lucide

import (
	html "github.com/plainkit/html"
)

// Undo creates a Undo Lucide icon.
func Undo(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-undo", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 7v6h6")),
		html.SvgPath(html.AD("M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
