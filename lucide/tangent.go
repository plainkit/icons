package lucide

import (
	html "github.com/plainkit/html"
)

// Tangent creates a Tangent Lucide icon.
func Tangent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tangent", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("4"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M15.59 5.41 5.41 15.59"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("17"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M12 22s-4-9-1.5-11.5S22 12 22 12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
