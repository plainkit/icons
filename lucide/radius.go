package lucide

import (
	html "github.com/plainkit/html"
)

// Radius creates a Radius Lucide icon.
func Radius(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radius", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20.34 17.52a10 10 0 1 0-2.82 2.82")),
		html.SvgCircle(html.ACx("19"), html.ACy("19"), html.AR("2")),
		html.SvgPath(html.AD("m13.41 13.41 4.18 4.18")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
