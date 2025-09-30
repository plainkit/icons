package lucide

import (
	html "github.com/plainkit/html"
)

// GripVertical creates a Grip Vertical Lucide icon.
func GripVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grip-vertical", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("9"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("9"), html.ACy("5"), html.AR("1")),
		html.SvgCircle(html.ACx("9"), html.ACy("19"), html.AR("1")),
		html.SvgCircle(html.ACx("15"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("15"), html.ACy("5"), html.AR("1")),
		html.SvgCircle(html.ACx("15"), html.ACy("19"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
