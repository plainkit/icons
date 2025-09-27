package lucide

import (
	html "github.com/plainkit/html"
)

// Codepen creates a Codepen Lucide icon.
func Codepen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-codepen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPolygon(html.APoints("12 2 22 8.5 22 15.5 12 22 2 15.5 2 8.5 12 2"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("22"), html.AY2("15.5"))),
		html.Child(html.SvgPolyline(html.APoints("22 8.5 12 15.5 2 8.5"))),
		html.Child(html.SvgPolyline(html.APoints("2 15.5 12 8.5 22 15.5"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("2"), html.AY2("8.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
