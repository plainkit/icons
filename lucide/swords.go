package lucide

import (
	html "github.com/plainkit/html"
)

// Swords creates a Swords Lucide icon.
func Swords(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-swords", args)
	children := []html.SvgArg{
		html.Child(html.SvgPolyline(html.APoints("14.5 17.5 3 6 3 3 6 3 17.5 14.5"))),
		html.Child(html.SvgLine(html.AX1("13"), html.AX2("19"), html.AY1("19"), html.AY2("13"))),
		html.Child(html.SvgLine(html.AX1("16"), html.AX2("20"), html.AY1("16"), html.AY2("20"))),
		html.Child(html.SvgLine(html.AX1("19"), html.AX2("21"), html.AY1("21"), html.AY2("19"))),
		html.Child(html.SvgPolyline(html.APoints("14.5 6.5 18 3 21 3 21 6 17.5 9.5"))),
		html.Child(html.SvgLine(html.AX1("5"), html.AX2("9"), html.AY1("14"), html.AY2("18"))),
		html.Child(html.SvgLine(html.AX1("7"), html.AX2("4"), html.AY1("17"), html.AY2("20"))),
		html.Child(html.SvgLine(html.AX1("3"), html.AX2("5"), html.AY1("19"), html.AY2("21"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
