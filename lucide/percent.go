package lucide

import (
	html "github.com/plainkit/html"
)

// Percent creates a Percent Lucide icon.
func Percent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-percent", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("19"), html.AX2("5"), html.AY1("5"), html.AY2("19")),
		html.SvgCircle(html.ACx("6.5"), html.ACy("6.5"), html.AR("2.5")),
		html.SvgCircle(html.ACx("17.5"), html.ACy("17.5"), html.AR("2.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
