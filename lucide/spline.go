package lucide

import (
	html "github.com/plainkit/html"
)

// Spline creates a Spline Lucide icon.
func Spline(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spline", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("19"), html.ACy("5"), html.AR("2")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("2")),
		html.SvgPath(html.AD("M5 17A12 12 0 0 1 17 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
