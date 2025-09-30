package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePause creates a Square Pause Lucide icon.
func SquarePause(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-pause", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgLine(html.AX1("10"), html.AX2("10"), html.AY1("15"), html.AY2("9")),
		html.SvgLine(html.AX1("14"), html.AX2("14"), html.AY1("15"), html.AY2("9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
