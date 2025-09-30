package lucide

import (
	html "github.com/plainkit/html"
)

// Palette creates a Palette Lucide icon.
func Palette(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-palette", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22a1 1 0 0 1 0-20 10 9 0 0 1 10 9 5 5 0 0 1-5 5h-2.25a1.75 1.75 0 0 0-1.4 2.8l.3.4a1.75 1.75 0 0 1-1.4 2.8z")),
		html.SvgCircle(html.ACx("13.5"), html.ACy("6.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgCircle(html.ACx("17.5"), html.ACy("10.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgCircle(html.ACx("6.5"), html.ACy("12.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgCircle(html.ACx("8.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
