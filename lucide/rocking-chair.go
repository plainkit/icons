package lucide

import (
	html "github.com/plainkit/html"
)

// RockingChair creates a Rocking Chair Lucide icon.
func RockingChair(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rocking-chair", args)
	children := []html.SvgArg{
		html.Child(html.SvgPolyline(html.APoints("3.5 2 6.5 12.5 18 12.5"))),
		html.Child(html.SvgLine(html.AX1("9.5"), html.AX2("5.5"), html.AY1("12.5"), html.AY2("20"))),
		html.Child(html.SvgLine(html.AX1("15"), html.AX2("18.5"), html.AY1("12.5"), html.AY2("20"))),
		html.Child(html.SvgPath(html.AD("M2.75 18a13 13 0 0 0 18.5 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
