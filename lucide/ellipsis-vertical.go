package lucide

import (
	html "github.com/plainkit/html"
)

// EllipsisVertical creates a Ellipsis Vertical Lucide icon.
func EllipsisVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ellipsis-vertical", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("1")),
		html.SvgCircle(html.ACx("12"), html.ACy("19"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
