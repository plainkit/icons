package lucide

import (
	html "github.com/plainkit/html"
)

// Scale3d creates a Scale 3d Lucide icon.
func Scale3d(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scale-3d", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 7v11a1 1 0 0 0 1 1h11")),
		html.SvgPath(html.AD("M5.293 18.707 11 13")),
		html.SvgCircle(html.ACx("19"), html.ACy("19"), html.AR("2")),
		html.SvgCircle(html.ACx("5"), html.ACy("5"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
