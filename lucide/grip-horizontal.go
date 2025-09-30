package lucide

import (
	html "github.com/plainkit/html"
)

// GripHorizontal creates a Grip Horizontal Lucide icon.
func GripHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grip-horizontal", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("9"), html.AR("1")),
		html.SvgCircle(html.ACx("19"), html.ACy("9"), html.AR("1")),
		html.SvgCircle(html.ACx("5"), html.ACy("9"), html.AR("1")),
		html.SvgCircle(html.ACx("12"), html.ACy("15"), html.AR("1")),
		html.SvgCircle(html.ACx("19"), html.ACy("15"), html.AR("1")),
		html.SvgCircle(html.ACx("5"), html.ACy("15"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
