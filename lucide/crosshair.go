package lucide

import (
	html "github.com/plainkit/html"
)

// Crosshair creates a Crosshair Lucide icon.
func Crosshair(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-crosshair", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgLine(html.AX1("22"), html.AX2("18"), html.AY1("12"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("2"), html.AY1("12"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("6"), html.AY2("2"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("22"), html.AY2("18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
