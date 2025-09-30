package lucide

import (
	html "github.com/plainkit/html"
)

// Sword creates a Sword Lucide icon.
func Sword(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sword", args)
	children := []html.SvgArg{
		html.SvgPolyline(html.APoints("14.5 17.5 3 6 3 3 6 3 17.5 14.5")),
		html.SvgLine(html.AX1("13"), html.AX2("19"), html.AY1("19"), html.AY2("13")),
		html.SvgLine(html.AX1("16"), html.AX2("20"), html.AY1("16"), html.AY2("20")),
		html.SvgLine(html.AX1("19"), html.AX2("21"), html.AY1("21"), html.AY2("19")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
