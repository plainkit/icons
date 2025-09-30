package lucide

import (
	html "github.com/plainkit/html"
)

// RectangleCircle creates a Rectangle Circle Lucide icon.
func RectangleCircle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rectangle-circle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 4v16H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1z")),
		html.SvgCircle(html.ACx("14"), html.ACy("12"), html.AR("8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
