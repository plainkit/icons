package lucide

import (
	html "github.com/plainkit/html"
)

// Diameter creates a Diameter Lucide icon.
func Diameter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-diameter", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("19"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("5"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M6.48 3.66a10 10 0 0 1 13.86 13.86"))),
		html.Child(html.SvgPath(html.AD("m6.41 6.41 11.18 11.18"))),
		html.Child(html.SvgPath(html.AD("M3.66 6.48a10 10 0 0 0 13.86 13.86"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
