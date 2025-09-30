package lucide

import (
	html "github.com/plainkit/html"
)

// Ellipsis creates a Ellipsis Lucide icon.
func Ellipsis(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ellipsis", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("19"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("5"), html.ACy("12"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
