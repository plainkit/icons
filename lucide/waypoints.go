package lucide

import (
	html "github.com/plainkit/html"
)

// Waypoints creates a Waypoints Lucide icon.
func Waypoints(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-waypoints", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("4.5"), html.AR("2.5"))),
		html.Child(html.SvgPath(html.AD("m10.2 6.3-3.9 3.9"))),
		html.Child(html.SvgCircle(html.ACx("4.5"), html.ACy("12"), html.AR("2.5"))),
		html.Child(html.SvgPath(html.AD("M7 12h10"))),
		html.Child(html.SvgCircle(html.ACx("19.5"), html.ACy("12"), html.AR("2.5"))),
		html.Child(html.SvgPath(html.AD("m13.8 17.7 3.9-3.9"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("19.5"), html.AR("2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
