package lucide

import (
	html "github.com/plainkit/html"
)

// Dessert creates a Dessert Lucide icon.
func Dessert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dessert", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.162 3.167A10 10 0 0 0 2 13a2 2 0 0 0 4 0v-1a2 2 0 0 1 4 0v4a2 2 0 0 0 4 0v-4a2 2 0 0 1 4 0v1a2 2 0 0 0 4-.006 10 10 0 0 0-8.161-9.826"))),
		html.Child(html.SvgPath(html.AD("M20.804 14.869a9 9 0 0 1-17.608 0"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("4"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
