package lucide

import (
	html "github.com/plainkit/html"
)

// Ambulance creates a Ambulance Lucide icon.
func Ambulance(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ambulance", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10H6"))),
		html.Child(html.SvgPath(html.AD("M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2"))),
		html.Child(html.SvgPath(html.AD("M19 18h2a1 1 0 0 0 1-1v-3.28a1 1 0 0 0-.684-.948l-1.923-.641a1 1 0 0 1-.578-.502l-1.539-3.076A1 1 0 0 0 16.382 8H14"))),
		html.Child(html.SvgPath(html.AD("M8 8v4"))),
		html.Child(html.SvgPath(html.AD("M9 18h6"))),
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("18"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("7"), html.ACy("18"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
