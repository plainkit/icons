package lucide

import (
	html "github.com/plainkit/html"
)

// Sheet creates a Sheet Lucide icon.
func Sheet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sheet", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2")),
		html.SvgLine(html.AX1("3"), html.AX2("21"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("3"), html.AX2("21"), html.AY1("15"), html.AY2("15")),
		html.SvgLine(html.AX1("9"), html.AX2("9"), html.AY1("9"), html.AY2("21")),
		html.SvgLine(html.AX1("15"), html.AX2("15"), html.AY1("9"), html.AY2("21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
