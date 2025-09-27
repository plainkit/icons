package lucide

import (
	html "github.com/plainkit/html"
)

// UndoDot creates a Undo Dot Lucide icon.
func UndoDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-undo-dot", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 17a9 9 0 0 0-15-6.7L3 13"))),
		html.Child(html.SvgPath(html.AD("M3 7v6h6"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("17"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
