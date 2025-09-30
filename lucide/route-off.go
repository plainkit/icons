package lucide

import (
	html "github.com/plainkit/html"
)

// RouteOff creates a Route Off Lucide icon.
func RouteOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-route-off", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("6"), html.ACy("19"), html.AR("3")),
		html.SvgPath(html.AD("M9 19h8.5c.4 0 .9-.1 1.3-.2")),
		html.SvgPath(html.AD("M5.2 5.2A3.5 3.53 0 0 0 6.5 12H12")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M21 15.3a3.5 3.5 0 0 0-3.3-3.3")),
		html.SvgPath(html.AD("M15 5h-4.3")),
		html.SvgCircle(html.ACx("18"), html.ACy("5"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
