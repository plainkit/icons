package lucide

import (
	html "github.com/plainkit/html"
)

// Brush creates a Brush Lucide icon.
func Brush(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brush", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m11 10 3 3"))),
		html.Child(html.SvgPath(html.AD("M6.5 21A3.5 3.5 0 1 0 3 17.5a2.62 2.62 0 0 1-.708 1.792A1 1 0 0 0 3 21z"))),
		html.Child(html.SvgPath(html.AD("M9.969 17.031 21.378 5.624a1 1 0 0 0-3.002-3.002L6.967 14.031"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
