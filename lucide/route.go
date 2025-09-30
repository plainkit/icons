package lucide

import (
	html "github.com/plainkit/html"
)

// Route creates a Route Lucide icon.
func Route(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-route", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("6"), html.ACy("19"), html.AR("3")),
		html.SvgPath(html.AD("M9 19h8.5a3.5 3.5 0 0 0 0-7h-11a3.5 3.5 0 0 1 0-7H15")),
		html.SvgCircle(html.ACx("18"), html.ACy("5"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
