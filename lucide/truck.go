package lucide

import (
	html "github.com/plainkit/html"
)

// Truck creates a Truck Lucide icon.
func Truck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-truck", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2")),
		html.SvgPath(html.AD("M15 18H9")),
		html.SvgPath(html.AD("M19 18h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.624l-3.48-4.35A1 1 0 0 0 17.52 8H14")),
		html.SvgCircle(html.ACx("17"), html.ACy("18"), html.AR("2")),
		html.SvgCircle(html.ACx("7"), html.ACy("18"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
