package lucide

import (
	html "github.com/plainkit/html"
)

// Grip creates a Grip Lucide icon.
func Grip(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grip", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("5"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("5"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("12"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("12"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("19"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("19"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
