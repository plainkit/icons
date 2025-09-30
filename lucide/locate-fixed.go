package lucide

import (
	html "github.com/plainkit/html"
)

// LocateFixed creates a Locate Fixed Lucide icon.
func LocateFixed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-locate-fixed", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("2"), html.AX2("5"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("19"), html.AX2("22"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("2"), html.AY2("5")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("19"), html.AY2("22")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("7")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
