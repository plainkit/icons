package lucide

import (
	html "github.com/plainkit/html"
)

// Currency creates a Currency Lucide icon.
func Currency(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-currency", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("8")),
		html.SvgLine(html.AX1("3"), html.AX2("6"), html.AY1("3"), html.AY2("6")),
		html.SvgLine(html.AX1("21"), html.AX2("18"), html.AY1("3"), html.AY2("6")),
		html.SvgLine(html.AX1("3"), html.AX2("6"), html.AY1("21"), html.AY2("18")),
		html.SvgLine(html.AX1("21"), html.AX2("18"), html.AY1("21"), html.AY2("18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
